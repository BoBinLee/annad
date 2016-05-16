// Package profile provides implementations to generate CLG profiles.
package profile

import (
	"reflect"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGProfile represents the object type of the CLG profile object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGProfile spec.ObjectType = "clg-profile"
)

// Config represents the configuration used to create a new CLG profile object.
type Config struct {
	// Settings.

	// Body represents the CLG's implemented method body.
	Body string `json:"body,omitempty"`

	// HasChanged describes whether the CLG changed. A change might be a
	// renaming, a signature modification or even a reimplementation or bugfix.
	HasChanged bool `json:"hash_changed,omitempty"`

	// Hash represents the hashed value of the CLG's implemented method.
	Hash string `json:"hash,omitempty"`

	// Inputs represents the CLG's implemented method input parameter types.
	Inputs []string `json:"inputs,omitempty"`

	// Name represents the CLG's implemented method name.
	Name string `json:"name,omitempty"`

	// Outputs represents the CLG's implemented method output parameter types.
	Outputs []string `json:"outputs,omitempty"`
}

// DefaultConfig provides a default configuration to create a new CLG profile
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Body:       "",
		HasChanged: false,
		Hash:       "",
		Inputs:     nil,
		Name:       "",
		Outputs:    nil,
	}

	return newConfig
}

// New creates a new configured CLG profile object.
func New(config Config) (spec.CLGProfile, error) {
	newProfile := &profile{
		Config: config,

		CreatedAt: time.Now(),
		ID:        id.NewObjectID(id.Hex128),
		Mutex:     sync.Mutex{},
		Type:      ObjectTypeCLGProfile,
	}

	if newProfile.Body == "" {
		return nil, maskAnyf(invalidConfigError, "method body of CLG profile must not be empty")
	}
	if newProfile.Hash == "" {
		return nil, maskAnyf(invalidConfigError, "hash of CLG profile must not be empty")
	}
	if newProfile.Name == "" {
		return nil, maskAnyf(invalidConfigError, "method name of CLG profile must not be empty")
	}
	if len(newProfile.Outputs) == 0 {
		return nil, maskAnyf(invalidConfigError, "output types of CLG profile must not be empty")
	}

	return newProfile, nil
}

// NewEmptyProfile simply returns an empty, maybe invalid, profile object. This
// should only be used for things like unmarshaling.
func NewEmptyProfile() spec.CLGProfile {
	return &profile{}
}

type profile struct {
	Config

	// CreatedAt represents the creation time of the profile.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// ID represents the profile's identifier.
	ID spec.ObjectID `json:"id,omitempty"`

	Mutex sync.Mutex `json:"-"`

	// Type represents the profile's object type.
	Type spec.ObjectType `json:"type,omitempty"`
}

func (p *profile) Equals(other spec.CLGProfile) bool {
	if p.GetBody() != other.GetBody() {
		return false
	}
	if p.GetHash() != other.GetHash() {
		return false
	}
	if !reflect.DeepEqual(p.GetInputs(), other.GetInputs()) {
		return false
	}
	if p.GetName() != other.GetName() {
		return false
	}
	if !reflect.DeepEqual(p.GetOutputs(), other.GetOutputs()) {
		return false
	}

	return true
}

func (p *profile) GetBody() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Body
}

func (p *profile) GetHasChanged() bool {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.HasChanged
}

func (p *profile) GetHash() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Hash
}

func (p *profile) GetInputs() []string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Inputs
}

func (p *profile) GetName() string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Name
}

func (p *profile) GetOutputs() []string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.Outputs
}

func (p *profile) SetHashChanged(hasChanged bool) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.HasChanged = hasChanged
}

func (p *profile) SetID(id spec.ObjectID) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.ID = id
}
package object

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type Note struct {
	// Hash of the tag.
	Hash plumbing.Hash
	// Message is an arbitrary text message.
	Message string
	// TargetType is the object type of the target.
	TargetType plumbing.ObjectType
	// Target is the hash of the target object.
	Target plumbing.Hash

	s storer.EncodedObjectStorer
}

// GetNote gets a note from an object storer and decodes it.
func GetNote(s storer.EncodedObjectStorer, h plumbing.Hash) (*Note, error) {
	o, err := s.EncodedObject(plumbing.TagObject, h)
	if err != nil {
		return nil, err
	}

	return DecodeNote(s, o)
}

func DecodeNote(s storer.EncodedObjectStorer, o plumbing.EncodedObject) (*Note, error) {
	n := &Note{s: s}
	if err := n.Decode(o); err != nil {
		return nil, err
	}

	return n, nil
}

// ID returns the object ID of the note. The returned value will always match
// the current value of Note.Hash.
//
// ID is present to fulfill the Object interface.
func (n *Note) ID() plumbing.Hash {
	return n.Hash
}

// Type returns the type of object. It always returns plumbing.NoteObject.
//
// Type is present to fulfill the Object interface.
func (n *Note) Type() plumbing.ObjectType {
	return plumbing.NoteObject
}

// Decode transforms a plumbing.EncodedObject into a Note struct.
func (n *Note) Decode(o plumbing.EncodedObject) (err error) {
	if o.Type() != plumbing.NoteObject {
		return ErrUnsupportedObject
	}

	n.Hash = o.Hash()

	return nil
}

// Encode transforms a Note into a plumbing.EncodedObject.
func (t *Note) Encode(o plumbing.EncodedObject) error {
	o.SetType(plumbing.NoteObject)

	return nil
}

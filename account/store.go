package account

import (
	"bytes"
	"encoding/gob"

	"github.com/frankh/nano/store"
)

// TODO: Decide whether accounts should be
// stored separately, or under wallets
type AccountStore struct {
	s *store.Store
}

func NewAccountStore(store *store.Store) *AccountStore {
	s := new(AccountStore)

	s.s = store

	return s
}

func (s *AccountStore) SetAccount(a *Account) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(a); err != nil {
		return err
	}

	return s.s.Set([]byte("account:"+a.Address()), buf.Bytes())
}

func (s *AccountStore) GetAccount(id string) (*Account, error) {
	v, err := s.s.Get([]byte("account:" + id))
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(v)
	dec := gob.NewDecoder(buf)

	var a *Account
	if err = dec.Decode(&a); err != nil {
		return nil, err
	}

	return a, nil
}
package wallet

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/frankh/crypto/ed25519"
	"github.com/frankh/nano/account"
	"github.com/frankh/nano/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Wallet struct {
	Id       string
	Seed     string
	Accounts map[string]*account.Account
}

func NewWallet() *Wallet {
	w := new(Wallet)

	w.Accounts = make(map[string]*account.Account)

	return w
}

func (w *Wallet) Init() (string, error) {
	if err := w.GenerateSeed(); err != nil {
		return "", errors.Wrap(err, "generating seed failed")
	}

	if err := w.GenerateID(); err != nil {
		return "", errors.Wrap(err, "generating id failed")
	}

	return w.Id, nil
}

func (w *Wallet) GenerateSeed() error {
	_, prv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}

	w.Seed = strings.ToUpper(hex.EncodeToString(prv))

	return nil
}

func (w *Wallet) GenerateID() error {
	pub, _, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}

	w.Id = strings.ToUpper(hex.EncodeToString(pub))

	return nil
}

func (w *Wallet) NewAccount() *account.Account {
	a := account.NewAccount()

	pub, prv, _ := types.KeypairFromSeed(w.Seed, uint32(len(w.Accounts)))
	a.PublicKey = pub
	a.PrivateKey = prv
	w.Accounts[string(a.Address())] = a

	return a
}

func (w *Wallet) String() string {
	b, err := json.Marshal(w)
	if err != nil {
		log.Warn(err)
		return w.Id
	}

	return string(b)
}

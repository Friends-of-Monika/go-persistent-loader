package persistent

import (
	"compress/zlib"
	"errors"
	"os"

	"github.com/nlpodyssey/gopickle/pickle"
	"mon.icu/go-persistent-loader/persistent/types"
)

func Load(path string) (*types.RenpyPersistentPersistent, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decomp, err := zlib.NewReader(file)
	if err != nil {
		_ = file.Close()
		return nil, err
	}

	unpk := pickle.NewUnpickler(decomp)
	unpk.FindClass = func(module, name string) (interface{}, error) {
		rt := types.GetRegisteredType(module, name)
		if rt == nil {
			return nil, errors.New(module + "." + name)
		}
		return rt, nil
	}

	data, err := unpk.Load()
	if err != nil {
		_ = decomp.Close()
		_ = file.Close()
		return nil, err
	}

	_ = decomp.Close()
	_ = file.Close()
	return data.(*types.RenpyPersistentPersistent), nil
}

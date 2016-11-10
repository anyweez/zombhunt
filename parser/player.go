package parser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anyweez/zombhunt/types"
	"github.com/anyweez/zombhunt/world"
)

const (
	BeltSize      = 8
	InventorySize = 32
)

func LoadInventory(filepath string) *types.Inventory {
	inventory := types.NewInventory()
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	// fast forward to belt section
	pos, err := ff_0x0800(data)

	for i := 0; i < BeltSize; i++ {
		next, err := readItem(data[pos+(i*15) : pos+(i*15)+15])

		if next.Id > 0 && err == nil {
			inventory.AddItem(next)
		}
	}

	pos += BeltSize*15 + 3 // 2 for inventory size, 1 for 0x03

	// read inventory
	for i := 0; i < InventorySize; i++ {
		next, err := readItem(data[pos+(i*15) : pos+(i*15)+15])

		if next.Id > 0 && err == nil {
			inventory.AddItem(next)
		}
	}

	return inventory
}

func ff_0x0800(data []byte) (int, error) {
	for i := 0; i < len(data)-2; i++ {
		if data[i] == 0x08 && data[i+1] == 0x00 {
			// Skip these two bytes and the 0x03 that comes after it.
			return i + 3, nil
		}
	}

	return -1, errors.New("No inventory section of player file detected")
}

func readItem(data []byte) (types.InventoryItem, error) {
	var item types.InventoryItem

	// The item ID is a two-byte value where the first half of the second byte is discarded
	// or potentially used for something non-ID related.
	first := uint32(binary.BigEndian.Uint16([]byte{0x00, data[0]}))
	second := uint32(binary.BigEndian.Uint16([]byte{data[1] & 15, 0x00}))
	item.Id = first + second

	// Read the item name from the world database
	target := world.Get().FindItem(item.Id)

	if target != nil {
		item.Name = world.Get().FindItem(item.Id).Name
	} else if item.Id == 0 {
		item.Name = "<empty>"
	} else {
		return item, fmt.Errorf("Unknown item: %d", item.Id)
	}

	// Read in the quantity from a set position
	item.Quantity = uint32(binary.BigEndian.Uint16([]byte{data[11], data[12]}))

	return item, nil
}

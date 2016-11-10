/**
 * Reads and stores all information related to the world (items, players, etc).
 */

package world

import "github.com/anyweez/zombhunt/types"

// Singleton World that can be loaded once then fetched multiple times.
var w types.World

func Get() *types.World {
	return &w
}

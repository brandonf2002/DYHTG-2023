package scenes

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gopxl/pixel"
	"golang.org/x/image/colornames"
)

type Pair struct {
	spider int
	candy  int
}
type Matching []*Pair

type SceneSpiders struct {
	game           *Game
	entityManager  EntityManager
	id             int
	background     pixel.Picture
	spiderPrefs    [][]int
	candyPrefs     [][]int
	playerMatching Matching
	xOffset        float64
	yOffset        float64
	size           float64
}

func NewSceneSpiders(game *Game) *SceneSpiders {
	ssp := SceneSpiders{game: game, id: 0, background: game.Assets.GetPicture("menu_background"), playerMatching: make(Matching, 5), xOffset: 100, yOffset: 100, size: 64}
	ssp.entityManager = make([]ComponentVector, 256)

	spiderSprites := make([]*pixel.Sprite, 5)
	candySprites := make([]*pixel.Sprite, 5)

	for i := 0; i < 5; i++ {
		spider_pic := game.Assets.GetPicture("spider" + strconv.Itoa(i+1))
		spiderSprites[i] = pixel.NewSprite(spider_pic, spider_pic.Bounds())
		candy_pic := game.Assets.GetPicture("candy" + strconv.Itoa(i+1))
		candySprites[i] = pixel.NewSprite(candy_pic, candy_pic.Bounds())

		spider := ssp.AddEntity()
		spider.Transform = NewCTransform(pixel.V(ssp.xOffset, ssp.yOffset+float64(i)*ssp.size), pixel.ZV, pixel.V(ssp.size/spider_pic.Bounds().W(), ssp.size/spider_pic.Bounds().H()), pixel.ZV, 0, 0)
		spider.BoundingBox = NewCBoundingBox(pixel.V(ssp.size, ssp.size))
		spider.Sprite = NewCSprite(spiderSprites[i])

		candy := ssp.AddEntity()
		candy.Transform = NewCTransform(pixel.V(ssp.xOffset+500, ssp.yOffset+float64(i)*ssp.size), pixel.ZV, pixel.V(ssp.size/spider_pic.Bounds().W(), ssp.size/spider_pic.Bounds().H()), pixel.ZV, 0, 0)
		candy.BoundingBox = NewCBoundingBox(pixel.V(ssp.size, ssp.size))
		candy.Sprite = NewCSprite(candySprites[i])
	}

	ssp.generatePreferences()

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			candy := ssp.AddEntity()
			candy.Transform = NewCTransform(pixel.V(ssp.xOffset+float64(j+1)*ssp.size, ssp.yOffset+float64(4-i)*ssp.size), pixel.ZV, pixel.V(ssp.size/candySprites[0].Frame().W(), ssp.size/candySprites[0].Frame().H()), pixel.ZV, 0, 0)
			candy.BoundingBox = NewCBoundingBox(pixel.V(ssp.size, ssp.size))
			candy.Sprite = NewCSprite(candySprites[ssp.spiderPrefs[4-i][j]])

			spider := ssp.AddEntity()
			spider.Transform = NewCTransform(pixel.V(ssp.xOffset+500+float64(j+1)*ssp.size, ssp.yOffset+float64(4-i)*ssp.size), pixel.ZV, pixel.V(ssp.size/candySprites[0].Frame().W(), ssp.size/candySprites[0].Frame().H()), pixel.ZV, 0, 0)
			spider.BoundingBox = NewCBoundingBox(pixel.V(ssp.size, ssp.size))
			spider.Sprite = NewCSprite(spiderSprites[ssp.candyPrefs[4-i][j]])
		}
	}

	return &ssp
}

func (ssp *SceneSpiders) generatePreferences() {
	ssp.spiderPrefs = make([][]int, 5)
	ssp.candyPrefs = make([][]int, 5)
	for i := 0; i < 5; i++ {
		ssp.spiderPrefs[i] = []int{0, 1, 2, 3, 4}
		ssp.candyPrefs[i] = []int{0, 1, 2, 3, 4}
		for j := 4; j > 0; j-- {
			r1 := rand.Intn(j + 1)
			r2 := rand.Intn(j + 1)
			ssp.spiderPrefs[i][r1], ssp.spiderPrefs[i][j] = ssp.spiderPrefs[i][j], ssp.spiderPrefs[i][r1]
			ssp.candyPrefs[i][r2], ssp.candyPrefs[i][j] = ssp.candyPrefs[i][j], ssp.candyPrefs[i][r2]
		}
	}
}

func (ssp *SceneSpiders) sLifespan() {
	for i, entity := range ssp.entityManager {
		if (CLifeSpan{}) != entity.LifeSpan {
			if entity.LifeSpan.FrameCounter >= entity.LifeSpan.NumOfFrames {
				ssp.entityManager[i].Sprite = CSprite{}
			}
			ssp.entityManager[i].LifeSpan.FrameCounter += 1
		}
	}
}

func (ssp *SceneSpiders) AddEntity() *ComponentVector {
	ssp.entityManager[ssp.id] = ComponentVector{}
	ssp.id += 1
	return &ssp.entityManager[ssp.id-1]
}

func (ssp *SceneSpiders) GetEntityManager() EntityManager {
	return ssp.entityManager
}

func (ssp *SceneSpiders) Update() {
	ssp.sLifespan()
	ssp.Render()
}

func (ssp *SceneSpiders) Render() {
	//sprite := pixel.NewSprite(ssp.background, ssp.background.Bounds())

	//scaleX := ssp.game.Window.Bounds().W() / ssp.background.Bounds().W()
	//scaleY := ssp.game.Window.Bounds().H() / ssp.background.Bounds().H()
	//sprite.Draw(ssp.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(ssp.game.Window.Bounds().Center()))

	ssp.game.Window.Clear(colornames.Antiquewhite)

	for _, entity := range ssp.entityManager {
		if (CTransform{}) != entity.Transform && (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(ssp.game.Window, pixel.IM.ScaledXY(pixel.ZV, entity.Transform.Scale).Rotated(pixel.ZV, entity.Transform.Angle).Moved(Add(entity.Transform.Pos, entity.BoundingBox.Half())))
		}
	}
}

func (ssp *SceneSpiders) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
		circleRemoved := false
		for i := 60; i < len(ssp.entityManager); i++ {
			if (CTag{}) != ssp.entityManager[i].Tag &&
				(CSprite{}) != ssp.entityManager[i].Sprite &&
				strings.HasPrefix(*ssp.entityManager[i].Tag.Tag, "red") &&
				Inside(action.Coords, ssp.entityManager[i]) {

				ssp.entityManager[i].Sprite = CSprite{}

				a, _ := strconv.Atoi((*ssp.entityManager[i].Tag.Tag)[3:4])
				b, _ := strconv.Atoi((*ssp.entityManager[i].Tag.Tag)[4:5])

				ssp.removePair(&Pair{a, b})

				// remove the corresponding circle
				for _, j := range []int{-1, 1} {
					if (CTag{}) != ssp.entityManager[i+j].Tag && *ssp.entityManager[i+j].Tag.Tag == *ssp.entityManager[i].Tag.Tag {
						ssp.entityManager[i+j].Sprite = CSprite{}
					}
				}

				circleRemoved = true
			}
		}
		if !circleRemoved {
			for i := 10; i < 60; i++ {
				if Inside(action.Coords, ssp.entityManager[i]) {
					row := 5 - ((i / 2) / 5)
					col := (i / 2) % 5
					newPair := &Pair{}
					addPair := false

					if i%2 == 0 && m(row, ssp.spiderPrefs[row][col], ssp.playerMatching) == -1 {
						newPair = &Pair{row, ssp.spiderPrefs[row][col]}
						addPair = true
					} else if i%2 == 1 && m(ssp.candyPrefs[row][col], row, ssp.playerMatching) == -1 {
						newPair = &Pair{ssp.candyPrefs[row][col], row}
						addPair = true
					}
					if addPair {
						ssp.addPair(newPair)
						ssp.addMatchingCircles(newPair, false)

						complete, stable, blockingPair := ssp.checkIfMatchingComplete()
						if complete && !stable {
							ssp.addMatchingCircles(blockingPair, true)
							for i := 60; i < len(ssp.entityManager); i++ {
								if (CTag{}) != ssp.entityManager[i].Tag && strings.HasPrefix(*ssp.entityManager[i].Tag.Tag, "red") {
									ssp.entityManager[i].LifeSpan = NewCLifeSpan(60)
								}
							}
							ssp.clearMatching()
						} else if stable {
							fmt.Println("You won!!!!!!!")
						}
					}
				}
			}
		}
	}
	if action.Name == "ESC" {
		ssp.game.ChangeScene("MENU", nil)
	}
}

func (ssp *SceneSpiders) addMatchingCircles(pair *Pair, blocking bool) {
	pic := ssp.game.Assets.GetPicture("redcircle")
	tag := "red" + strconv.Itoa(pair.spider) + strconv.Itoa(pair.candy)
	if blocking {
		pic = ssp.game.Assets.GetPicture("blocking")
		tag = "blocking"
	}

	circ1 := ssp.AddEntity()
	circ1.Transform = NewCTransform(pixel.V(ssp.xOffset+float64(rank(pair.spider, pair.candy, ssp.spiderPrefs)+1)*ssp.size, ssp.yOffset+float64(pair.spider)*ssp.size), pixel.ZV, pixel.V(0.5, 0.5), pixel.ZV, 0, 0)
	circ1.BoundingBox = NewCBoundingBox(pixel.V(ssp.size, ssp.size))
	circ1.Sprite = NewCSprite(pixel.NewSprite(pic, pic.Bounds()))
	circ1.Tag = NewCTag(&tag)

	circ2 := ssp.AddEntity()
	circ2.Transform = NewCTransform(pixel.V(ssp.xOffset+500+float64(rank(pair.candy, pair.spider, ssp.candyPrefs)+1)*ssp.size, ssp.yOffset+float64(pair.candy)*ssp.size), pixel.ZV, pixel.V(0.5, 0.5), pixel.ZV, 0, 0)
	circ2.BoundingBox = NewCBoundingBox(pixel.V(ssp.size, ssp.size))
	circ2.Sprite = NewCSprite(pixel.NewSprite(pic, pic.Bounds()))
	circ2.Tag = NewCTag(&tag)

	if blocking {
		circ1.LifeSpan = NewCLifeSpan(60)
		circ2.LifeSpan = NewCLifeSpan(60)
	}
}

// returns (complete, stable, (optional) blocking pair)
func (ssp *SceneSpiders) checkIfMatchingComplete() (bool, bool, *Pair) {
	for _, pair := range ssp.playerMatching {
		if pair == nil {
			return false, false, nil
		}
	}
	// search for a blocking pair
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			M := ssp.playerMatching
			apfs := ssp.spiderPrefs
			bpfs := ssp.candyPrefs
			if prefers(i, j, m(i, -1, M), apfs) && prefers(j, i, m(-1, j, M), bpfs) {
				return true, false, &Pair{i, j}
			}
		}
	}
	return true, true, nil
}

// adds a pair to the matching
func (ssp *SceneSpiders) addPair(pair *Pair) {
	for i := 0; i < 5; i++ {
		if ssp.playerMatching[i] == nil {
			ssp.playerMatching[i] = pair
			return
		}
	}
}

// removes a pair from the matching
func (ssp *SceneSpiders) removePair(pair *Pair) bool {
	for i := 0; i < 5; i++ {
		if ssp.playerMatching[i].spider == pair.spider && ssp.playerMatching[i].candy == pair.candy {
			ssp.playerMatching[i] = nil
			return true
		}
	}
	return false
}

// removes all pairs from the matching
func (ssp *SceneSpiders) clearMatching() {
	for i := 0; i < 5; i++ {
		ssp.playerMatching[i] = nil
	}
}

// return matched object of either a a spider or b a candy
func m(a int, b int, matching Matching) int {
	for _, pair := range matching {
		if pair != nil && pair.spider == a {
			return pair.candy
		}
		if pair != nil && pair.candy == b {
			return pair.spider
		}
	}
	return -1
}

// returns the rank given to b by a according to prefs
func rank(a int, b int, prefs [][]int) int {
	for i, x := range prefs[a] {
		if b == x {
			return i
		}
	}
	return -1
}

// check whether a prefers b1 to b2 acording to prefs
func prefers(a int, b1 int, b2 int, prefs [][]int) bool {
	for _, b := range prefs[a] {
		if b == b2 {
			return false
		} else if b == b1 {
			return true
		}
	}
	return false
}

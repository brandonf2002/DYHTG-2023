package scenes

import (
	"fmt"
	"strings"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/text"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type SceneCodingChallenge struct {
	game          *Game
	entityManager EntityManager
	id            int
	background    pixel.Picture
	lets_go       *interp.Interpreter
}

func NewSceneCodingChallenge(game *Game) *SceneCodingChallenge {
	scc := SceneCodingChallenge{game: game, id: 0}
	scc.entityManager = make([]ComponentVector, 64)

	user_input := text.New(pixel.V(100, 750), scc.game.Assets.GetFont("basic"))

	term := scc.AddEntity()
	term.Text = NewCText(user_input)

	display_text := text.New(pixel.V(50, 800), scc.game.Assets.GetFont("basic"))
	// fmt.Fprintln(display_text, `Merge Two Sorted Lists

	// You are given the heads of two sorted linked lists list1 and list2.
	// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
	// Return the head of the merged linked list.`)
	wrapText(display_text, `Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2. Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists. Return the head of the merged linked list.`, 450)

	problem_statement := scc.AddEntity()
	problem_statement.Text = NewCText(display_text)

	scc.lets_go = interp.New(interp.Options{})
	scc.lets_go.Use(stdlib.Symbols)

	v, err := scc.lets_go.Eval(`
import (
	"fmt"
	"reflect"
)

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // dummy node to start the merged list
    dummy := &ListNode{}
    current := dummy

    // while both lists are not empty
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }

    // if list1 is not exhausted
    if list1 != nil {
        current.Next = list1
    }

    // if list2 is not exhausted
    if list2 != nil {
        current.Next = list2
    }

    return dummy.Next
}

// Utility function to build a linked list from a slice.
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Utility function to convert a linked list back to a slice.
func listToSlice(node *ListNode) []int {
	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	return nums
}

func eval_tests() bool {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, test := range tests {
		l1 := sliceToList(test.list1)
		l2 := sliceToList(test.list2)
		merged := mergeTwoLists(l1, l2)
		result := listToSlice(merged)
		if !(reflect.DeepEqual(result, test.expected)) {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, test := range tests {
		l1 := sliceToList(test.list1)
		l2 := sliceToList(test.list2)
		merged := mergeTwoLists(l1, l2)
		result := listToSlice(merged)
		if reflect.DeepEqual(result, test.expected) {
			fmt.Println("Passed!")
		} else {
			fmt.Printf("Failed! Expected %v but got %v\n", test.expected, result)
		}
	}
}
	`)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", v)

	v, err = scc.lets_go.Eval("eval_tests()")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", v)

	return &scc
}

// wrapText wraps the given text so that it fits within the given width
func wrapText(txt *text.Text, content string, maxWidth float64) {
	words := strings.Split(content, " ")
	var sb strings.Builder

	lineWidth := 0.0
	for _, word := range words {
		wordWidth := txt.BoundsOf(word + " ").W()
		if lineWidth+wordWidth > maxWidth {
			sb.WriteString("\n")
			lineWidth = 0
		}
		sb.WriteString(word + " ")
		lineWidth += wordWidth
	}
	txt.WriteString(sb.String())
}

func (scc *SceneCodingChallenge) AddEntity() *ComponentVector {
	scc.entityManager[scc.id] = ComponentVector{}
	scc.id += 1
	return &scc.entityManager[scc.id-1]
}

func (scc *SceneCodingChallenge) GetEntityManager() EntityManager {
	return scc.entityManager
}

func (scc *SceneCodingChallenge) Update() {

	scc.Render()
}

func (scc *SceneCodingChallenge) Render() {

	scc.game.Window.Clear(pixel.RGB(0, 0, 0))

	// display_text := text.New(pixel.V(100, 500), scc.game.Assets.GetFont("basic"))

	// display_text.Draw(scc.game.Window, pixel.IM)

	// sprite := pixel.NewSprite(scc.background, scc.background.Bounds())

	// scaleX := scc.game.Window.Bounds().W() / scc.background.Bounds().W()
	// scaleY := scc.game.Window.Bounds().H() / scc.background.Bounds().H()

	// sprite.Draw(scc.game.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(scaleX, scaleY)).Moved(scc.game.Window.Bounds().Center()))

	scc.entityManager[0].Text.Text.WriteString(scc.game.Window.Typed())

	for _, entity := range scc.entityManager {
		if (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(scc.game.Window, pixel.IM.Moved(entity.BoundingBox.Half()))
		}

		if (CText{}) != entity.Text {
			entity.Text.Text.Draw(scc.game.Window, pixel.IM.Scaled(entity.Text.Text.Orig, 2))
		}

	}
}

func (scc *SceneCodingChallenge) DoAction(action Action) {
	if action.Name == "LEFT_MOUSE" {
	}

	if action.Name == "ESC" {
		scc.game.ChangeScene("MENU", nil)
	}
}

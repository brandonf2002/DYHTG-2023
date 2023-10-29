package scenes

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
	"strings"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
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

	user_input := text.New(pixel.V(50, 700), scc.game.Assets.GetFont("basic"))
	user_input.WriteString(`/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */true
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
}`)

	term := scc.AddEntity()
	term.Text = NewCText(user_input)
	term.Text.Str_of_Text = `/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
}`
	term.Text.CursorPos = 183

	display_text := text.New(pixel.V(50, 800), scc.game.Assets.GetFont("basic"))
	wrapText(display_text, `Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2. Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists. Return the head of the merged linked list.`, 450)

	problem_statement := scc.AddEntity()
	problem_statement.Text = NewCText(display_text)

	return &scc
}

func eval_function(scc *SceneCodingChallenge, function string) bool {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	v, err := i.Eval(`
import (
	"fmt"
	"reflect"
)

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     // dummy node to start the merged list
//     dummy := &ListNode{}
//     current := dummy

//     // while both lists are not empty
//     for list1 != nil && list2 != nil {
//         if list1.Val < list2.Val {
//             current.Next = list1
//             list1 = list1.Next
//         } else {
//             current.Next = list2
//             list2 = list2.Next
//         }
//         current = current.Next
//     }

//     // if list1 is not exhausted
//     if list1 != nil {
//         current.Next = list1
//     }

//     // if list2 is not exhausted
//     if list2 != nil {
//         current.Next = list2
//     }

//     return dummy.Next
// }

// Utility function to build a linked list from a slice.
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	print("Hello")
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

	ret := true

	for _, test := range tests {
		l1 := sliceToList(test.list1)
		l2 := sliceToList(test.list2)
		merged := mergeTwoLists(l1, l2)
		result := listToSlice(merged)
		if !(reflect.DeepEqual(result, test.expected)) {
			ret = false
		}
		if reflect.DeepEqual(result, test.expected) {
			fmt.Println("Passed!")
		} else {
			fmt.Printf("Failed! Expected %v but got %v\n", test.expected, result)
		}
	}
	return ret
}

// func main() {
// 	// Test cases
// 	tests := []struct {
// 		list1    []int
// 		list2    []int
// 		expected []int
// 	}{
// 		{
// 			list1:    []int{1, 2, 4},
// 			list2:    []int{1, 3, 4},
// 			expected: []int{1, 1, 2, 3, 4, 4},
// 		},
// 		{
// 			list1:    []int{},
// 			list2:    []int{0},
// 			expected: []int{0},
// 		},
// 	}

// 	for _, test := range tests {
// 		l1 := sliceToList(test.list1)
// 		l2 := sliceToList(test.list2)
// 		merged := mergeTwoLists(l1, l2)
// 		result := listToSlice(merged)
// 		if reflect.DeepEqual(result, test.expected) {
// 			fmt.Println("Passed!")
// 		} else {
// 			fmt.Printf("Failed! Expected %v but got %v\n", test.expected, result)
// 		}
// 	}
// }
	` + "\n\n" + function)

	v, err = i.Eval("eval_tests()")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Result type 1: %v\n", v)

	fmt.Printf("Result type 2: %v\n", v.IsValid())
	if v.IsValid() && v.Type() == reflect.TypeOf(true) {
		println(v.Bool())
		return v.Bool()
	}

	println("Did not")
	return false
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

// Create a single-colored texture for the cursor
func createCursorTexture(scc SceneCodingChallenge) *pixel.PictureData {
	img := image.NewRGBA(image.Rect(0, 0, 3, int(scc.entityManager[0].Text.Text.LineHeight)*2))
	col := color.RGBA{0, 255, 0, 255} // Black color for the cursor
	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(int(x), int(y), col)
		}
	}
	return pixel.PictureDataFromImage(img)
}

func drawCursor(win *pixelgl.Window, txt *text.Text, cursorPos int, str string, scc SceneCodingChallenge) (float64, float64, *pixel.Sprite) {
	lines := strings.Split(str, "\n")

	totalChars := 0
	cursorPosX := txt.Orig.X
	cursorPosY := txt.Orig.Y
	lineHeight := txt.LineHeight

	for _, line := range lines {
		if totalChars+len(line) >= cursorPos {
			// If the cursor is within this line, calculate its X position
			cursorPosX += txt.BoundsOf(line[:cursorPos-totalChars]).W() * 2
			break
		} else {
			// Otherwise, move the Y position up by one line and continue
			cursorPosY -= lineHeight * 2
			totalChars += len(line) + 1 // +1 to account for the newline character
		}
	}

	cursorTexture := createCursorTexture(scc)
	cursorSprite := pixel.NewSprite(pixel.Picture(cursorTexture), cursorTexture.Bounds())
	return cursorPosX, cursorPosY, cursorSprite
}

func (scc *SceneCodingChallenge) Render() {

	scc.game.Window.Clear(pixel.RGB(0, 0, 0))
	char := scc.game.Window.Typed()
	if len(char) > 0 {
		// Insert char at the cursor's position
		scc.entityManager[0].Text.Str_of_Text = scc.entityManager[0].Text.Str_of_Text[:scc.entityManager[0].Text.CursorPos] + char + scc.entityManager[0].Text.Str_of_Text[scc.entityManager[0].Text.CursorPos:]
		scc.entityManager[0].Text.CursorPos += len(char)
		scc.entityManager[0].Text.Text.Clear()
		scc.entityManager[0].Text.Text.WriteString(scc.entityManager[0].Text.Str_of_Text)
	}

	// For the enter key:
	if scc.game.Window.JustPressed(pixelgl.KeyEnter) || scc.game.Window.Repeated(pixelgl.KeyEnter) {
		// Insert newline at the cursor's position
		scc.entityManager[0].Text.Str_of_Text = scc.entityManager[0].Text.Str_of_Text[:scc.entityManager[0].Text.CursorPos] + "\n" + scc.entityManager[0].Text.Str_of_Text[scc.entityManager[0].Text.CursorPos:]
		scc.entityManager[0].Text.CursorPos++ // Move cursor after the newline
		scc.entityManager[0].Text.Text.Clear()
		scc.entityManager[0].Text.Text.WriteString(scc.entityManager[0].Text.Str_of_Text)
	}

	// For the backspace key:
	if scc.game.Window.JustPressed(pixelgl.KeyBackspace) || scc.game.Window.Repeated(pixelgl.KeyBackspace) {
		if scc.entityManager[0].Text.CursorPos > 0 {
			// Remove the character before the cursor
			scc.entityManager[0].Text.Str_of_Text = scc.entityManager[0].Text.Str_of_Text[:scc.entityManager[0].Text.CursorPos-1] + scc.entityManager[0].Text.Str_of_Text[scc.entityManager[0].Text.CursorPos:]
			scc.entityManager[0].Text.CursorPos--
			scc.entityManager[0].Text.Text.Clear()
			scc.entityManager[0].Text.Text.WriteString(scc.entityManager[0].Text.Str_of_Text)
		}
	}

	// char := scc.game.Window.Typed()
	// scc.entityManager[0].Text.Text.WriteString(scc.game.Window.Typed())
	// scc.entityManager[0].Text.Str_of_Text += char

	if scc.game.Window.JustPressed(pixelgl.KeyLeft) || scc.game.Window.Repeated(pixelgl.KeyLeft) {
		if scc.entityManager[0].Text.CursorPos > 0 {
			scc.entityManager[0].Text.CursorPos--
		}
	} else if scc.game.Window.JustPressed(pixelgl.KeyRight) || scc.game.Window.Repeated(pixelgl.KeyRight) {
		if scc.entityManager[0].Text.CursorPos < len(scc.entityManager[0].Text.Str_of_Text) {
			scc.entityManager[0].Text.CursorPos++
		}
	}

	if scc.game.Window.JustPressed(pixelgl.KeyKPEnter) {
		println("Shift + Enter")
		eval_function(scc, scc.entityManager[0].Text.Str_of_Text)
	}

	num_lines := len(strings.Split(strings.Trim(scc.entityManager[0].Text.Str_of_Text, "\n"), "\n"))
	scc.entityManager[0].Text.Text.Draw(scc.game.Window, pixel.IM.Scaled(scc.entityManager[0].Text.Text.Orig.Add(pixel.V(0, float64(num_lines)*26)), 2))

	scc.entityManager[1].Text.Text.Draw(scc.game.Window, pixel.IM.Scaled(scc.entityManager[1].Text.Text.Orig, 2))

	num_lines = len(strings.Split(scc.entityManager[0].Text.Str_of_Text, "\n"))
	x, y, sprite := drawCursor(scc.game.Window, scc.entityManager[0].Text.Text, scc.entityManager[0].Text.CursorPos, scc.entityManager[0].Text.Str_of_Text, *scc)
	// sprite.Draw(scc.game.Window, pixel.IM.Moved(pos.Sub(pixel.V(0, float64(num_lines)*26))))
	sprite.Draw(scc.game.Window, pixel.IM.Moved(pixel.V(x+3, y-13)))

	for _, entity := range scc.entityManager {
		if (CBoundingBox{}) != entity.BoundingBox && (CSprite{}) != entity.Sprite {
			entity.Sprite.Sprite.Draw(scc.game.Window, pixel.IM.Moved(entity.BoundingBox.Half()))
		}

		if (CText{}) != entity.Text {

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

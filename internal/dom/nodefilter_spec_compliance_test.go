package dom

import (
	"testing"
)

// TestNodeFilterSpecCompliance tests compliance with WHATWG DOM Section 6.3
// Interface NodeFilter specification
func TestNodeFilterSpecCompliance(t *testing.T) {
	t.Run("WHATWG DOM Section 6.3 - Interface NodeFilter", func(t *testing.T) {
		t.Run("Constants for acceptNode() return values", func(t *testing.T) {
			// Verify spec-compliant constant values for filter return values
			if NodeFilterAccept != 1 {
				t.Errorf("Expected FILTER_ACCEPT to be 1, got %d", NodeFilterAccept)
			}

			if NodeFilterReject != 2 {
				t.Errorf("Expected FILTER_REJECT to be 2, got %d", NodeFilterReject)
			}

			if NodeFilterSkip != 3 {
				t.Errorf("Expected FILTER_SKIP to be 3, got %d", NodeFilterSkip)
			}
		})

		t.Run("Constants for whatToShow bitmask", func(t *testing.T) {
			// Test all whatToShow constants match spec values exactly
			testCases := []struct {
				name     string
				constant uint32
				expected uint32
				hex      string
				legacy   bool
			}{
				{"SHOW_ALL", NodeFilterShowAll, 0xFFFFFFFF, "FFFFFFFF", false},
				{"SHOW_ELEMENT", NodeFilterShowElement, 0x1, "1", false},
				{"SHOW_ATTRIBUTE", NodeFilterShowAttribute, 0x2, "2", true}, // legacy
				{"SHOW_TEXT", NodeFilterShowText, 0x4, "4", false},
				{"SHOW_CDATA_SECTION", NodeFilterShowCDataSection, 0x8, "8", false},
				{"SHOW_ENTITY_REFERENCE", NodeFilterShowEntityReference, 0x10, "10", true}, // legacy
				{"SHOW_ENTITY", NodeFilterShowEntity, 0x20, "20", true},                    // legacy
				{"SHOW_PROCESSING_INSTRUCTION", NodeFilterShowProcessingInstruction, 0x40, "40", false},
				{"SHOW_COMMENT", NodeFilterShowComment, 0x80, "80", false},
				{"SHOW_DOCUMENT", NodeFilterShowDocument, 0x100, "100", false},
				{"SHOW_DOCUMENT_TYPE", NodeFilterShowDocumentType, 0x200, "200", false},
				{"SHOW_DOCUMENT_FRAGMENT", NodeFilterShowDocumentFragment, 0x400, "400", false},
				{"SHOW_NOTATION", NodeFilterShowNotation, 0x800, "800", true}, // legacy
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					if tc.constant != tc.expected {
						t.Errorf("Expected %s to be 0x%s (%d), got 0x%x (%d)",
							tc.name, tc.hex, tc.expected, tc.constant, tc.constant)
					}

					// Verify decimal values match spec
					if tc.name == "SHOW_ALL" && tc.constant != 4294967295 {
						t.Errorf("Expected SHOW_ALL to be 4294967295 in decimal, got %d", tc.constant)
					}
					if tc.name == "SHOW_PROCESSING_INSTRUCTION" && tc.constant != 64 {
						t.Errorf("Expected SHOW_PROCESSING_INSTRUCTION to be 64 in decimal, got %d", tc.constant)
					}
					if tc.name == "SHOW_COMMENT" && tc.constant != 128 {
						t.Errorf("Expected SHOW_COMMENT to be 128 in decimal, got %d", tc.constant)
					}
					if tc.name == "SHOW_DOCUMENT" && tc.constant != 256 {
						t.Errorf("Expected SHOW_DOCUMENT to be 256 in decimal, got %d", tc.constant)
					}
					if tc.name == "SHOW_DOCUMENT_TYPE" && tc.constant != 512 {
						t.Errorf("Expected SHOW_DOCUMENT_TYPE to be 512 in decimal, got %d", tc.constant)
					}
					if tc.name == "SHOW_DOCUMENT_FRAGMENT" && tc.constant != 1024 {
						t.Errorf("Expected SHOW_DOCUMENT_FRAGMENT to be 1024 in decimal, got %d", tc.constant)
					}

					if tc.legacy {
						t.Logf("Note: %s is marked as legacy in the specification", tc.name)
					}
				})
			}
		})

		t.Run("Callback interface implementation", func(t *testing.T) {
			doc := NewDocument()
			elem := doc.CreateElement("div")
			text := doc.CreateTextNode("test")

			// Test that NodeFilter can be implemented as a function (callback interface)
			customFilter := NodeFilterFunc(func(node Node) int {
				if node.NodeType() == ElementNode {
					return NodeFilterAccept
				}
				return NodeFilterSkip
			})

			// Verify it implements the NodeFilter interface
			var filter NodeFilter = customFilter
			if filter == nil {
				t.Error("NodeFilterFunc should implement NodeFilter interface")
			}

			// Test acceptNode method behavior
			result := filter.AcceptNode(elem)
			if result != NodeFilterAccept {
				t.Errorf("Expected filter to accept element, got %d", result)
			}

			result = filter.AcceptNode(text)
			if result != NodeFilterSkip {
				t.Errorf("Expected filter to skip text node, got %d", result)
			}
		})

		t.Run("NodeFilter usage with NodeIterator", func(t *testing.T) {
			doc := NewDocument()
			root := doc.CreateElement("root")
			elem1 := doc.CreateElement("div")
			elem2 := doc.CreateElement("span")
			text1 := doc.CreateTextNode("text1")
			text2 := doc.CreateTextNode("text2")

			// Build structure: root -> [elem1 -> text1, elem2 -> text2]
			doc.AppendChild(root)
			root.AppendChild(elem1)
			root.AppendChild(elem2)
			elem1.AppendChild(text1)
			elem2.AppendChild(text2)

			// Test with whatToShow only (no custom filter)
			iterator := doc.CreateNodeIterator(root, NodeFilterShowElement, nil)

			// Should find root, elem1, elem2 (all elements)
			nodes := []Node{}
			for {
				node := iterator.NextNode()
				if node == nil {
					break
				}
				nodes = append(nodes, node)
			}

			expectedCount := 3 // root, elem1, elem2
			if len(nodes) != expectedCount {
				t.Errorf("Expected %d elements, got %d", expectedCount, len(nodes))
			}

			// Test with custom filter
			elementFilter := NodeFilterFunc(func(node Node) int {
				if elem, ok := node.(*Element); ok && elem.TagName() == "div" {
					return NodeFilterAccept
				}
				return NodeFilterSkip
			})

			iterator2 := doc.CreateNodeIterator(root, NodeFilterShowElement, elementFilter)

			// Should find only elem1 (the div)
			nodes2 := []Node{}
			for {
				node := iterator2.NextNode()
				if node == nil {
					break
				}
				nodes2 = append(nodes2, node)
			}

			if len(nodes2) != 1 {
				t.Errorf("Expected 1 div element, got %d", len(nodes2))
			}
			if len(nodes2) > 0 {
				if elem, ok := nodes2[0].(*Element); !ok || elem.TagName() != "div" {
					t.Errorf("Expected to find div element, got %v", nodes2[0])
				}
			}
		})

		t.Run("NodeFilter usage with TreeWalker", func(t *testing.T) {
			doc := NewDocument()
			root := doc.CreateElement("root")
			elem1 := doc.CreateElement("div")
			elem2 := doc.CreateElement("span")
			text1 := doc.CreateTextNode("text1")
			comment := doc.CreateComment("comment")

			// Build structure: root -> [elem1 -> text1, elem2, comment]
			doc.AppendChild(root)
			root.AppendChild(elem1)
			root.AppendChild(elem2)
			root.AppendChild(comment)
			elem1.AppendChild(text1)

			// Test with combined whatToShow (elements and comments)
			walker := doc.CreateTreeWalker(root, NodeFilterShowElement|NodeFilterShowComment, nil)

			// Should find root, elem1, elem2, comment
			nodes := []Node{}
			current := walker.CurrentNode()
			nodes = append(nodes, current) // Start with root

			for {
				node := walker.NextNode()
				if node == nil {
					break
				}
				nodes = append(nodes, node)
			}

			expectedCount := 4 // root, elem1, elem2, comment
			if len(nodes) != expectedCount {
				t.Errorf("Expected %d nodes (elements + comment), got %d", expectedCount, len(nodes))
			}

			// Verify types
			elementCount := 0
			commentCount := 0
			for _, node := range nodes {
				switch node.NodeType() {
				case ElementNode:
					elementCount++
				case CommentNode:
					commentCount++
				}
			}

			if elementCount != 3 { // root, elem1, elem2
				t.Errorf("Expected 3 elements, got %d", elementCount)
			}
			if commentCount != 1 {
				t.Errorf("Expected 1 comment, got %d", commentCount)
			}
		})

		t.Run("Filter return value behavior", func(t *testing.T) {
			doc := NewDocument()
			root := doc.CreateElement("root")
			elem1 := doc.CreateElement("accept")
			elem2 := doc.CreateElement("reject")
			elem3 := doc.CreateElement("skip")

			root.AppendChild(elem1)
			root.AppendChild(elem2)
			root.AppendChild(elem3)

			// Custom filter that returns different values based on tag name
			behaviorFilter := NodeFilterFunc(func(node Node) int {
				if elem, ok := node.(*Element); ok {
					switch elem.TagName() {
					case "accept":
						return NodeFilterAccept
					case "reject":
						return NodeFilterReject
					case "skip":
						return NodeFilterSkip
					}
				}
				return NodeFilterAccept // Default for root
			})

			iterator := doc.CreateNodeIterator(root, NodeFilterShowElement, behaviorFilter)

			// Should find root and "accept" element only
			// "reject" should be completely excluded from iteration
			// "skip" should be skipped but its children would be considered (none in this case)
			nodes := []Node{}
			for {
				node := iterator.NextNode()
				if node == nil {
					break
				}
				nodes = append(nodes, node)
			}

			expectedCount := 2 // root and "accept" element
			if len(nodes) != expectedCount {
				t.Errorf("Expected %d nodes (root + accept), got %d", expectedCount, len(nodes))
			}

			// Verify we got the right nodes
			if len(nodes) >= 2 {
				if elem, ok := nodes[1].(*Element); !ok || elem.TagName() != "accept" {
					t.Errorf("Expected second node to be 'accept' element, got %v", nodes[1])
				}
			}
		})
	})
}

func TestNodeFilterIntegrationCompliance(t *testing.T) {
	t.Run("Complex DOM structure filtering", func(t *testing.T) {
		doc := NewDocument()

		// Create a complex structure mimicking a real document
		html := doc.CreateElement("html")
		head := doc.CreateElement("head")
		title := doc.CreateElement("title")
		titleText := doc.CreateTextNode("Test Document")
		body := doc.CreateElement("body")
		header := doc.CreateElement("header")
		nav := doc.CreateElement("nav")
		main := doc.CreateElement("main")
		article := doc.CreateElement("article")
		section := doc.CreateElement("section")
		footer := doc.CreateElement("footer")
		comment1 := doc.CreateComment("Header comment")
		comment2 := doc.CreateComment("Footer comment")

		// Build the structure
		doc.AppendChild(html)
		html.AppendChild(head)
		html.AppendChild(body)
		head.AppendChild(title)
		title.AppendChild(titleText)
		body.AppendChild(header)
		body.AppendChild(comment1)
		body.AppendChild(nav)
		body.AppendChild(main)
		body.AppendChild(footer)
		body.AppendChild(comment2)
		main.AppendChild(article)
		article.AppendChild(section)

		// Test filtering for semantic HTML5 elements only
		semanticFilter := NodeFilterFunc(func(node Node) int {
			if elem, ok := node.(*Element); ok {
				semanticTags := map[string]bool{
					"header": true, "nav": true, "main": true,
					"article": true, "section": true, "footer": true,
				}
				if semanticTags[elem.TagName()] {
					return NodeFilterAccept
				}
			}
			return NodeFilterSkip
		})

		iterator := doc.CreateNodeIterator(html, NodeFilterShowElement, semanticFilter)

		semanticElements := []Node{}
		for {
			node := iterator.NextNode()
			if node == nil {
				break
			}
			semanticElements = append(semanticElements, node)
		}

		expectedSemanticCount := 6 // header, nav, main, article, section, footer
		if len(semanticElements) != expectedSemanticCount {
			t.Errorf("Expected %d semantic elements, got %d", expectedSemanticCount, len(semanticElements))
		}

		// Test filtering for text nodes and comments
		contentFilter := doc.CreateNodeIterator(html, NodeFilterShowText|NodeFilterShowComment, nil)

		contentNodes := []Node{}
		for {
			node := contentFilter.NextNode()
			if node == nil {
				break
			}
			contentNodes = append(contentNodes, node)
		}

		expectedContentCount := 3 // titleText, comment1, comment2
		if len(contentNodes) != expectedContentCount {
			t.Errorf("Expected %d content nodes, got %d", expectedContentCount, len(contentNodes))
		}
	})

	t.Run("Bitmask combination testing", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("root")
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("text")
		comment := doc.CreateComment("comment")

		root.AppendChild(elem)
		root.AppendChild(text)
		root.AppendChild(comment)

		testCases := []struct {
			name       string
			whatToShow uint32
			expected   []NodeType
		}{
			{
				"Elements only",
				NodeFilterShowElement,
				[]NodeType{ElementNode}, // only elem
			},
			{
				"Text only",
				NodeFilterShowText,
				[]NodeType{TextNode}, // only text
			},
			{
				"Comments only",
				NodeFilterShowComment,
				[]NodeType{CommentNode}, // only comment
			},
			{
				"Elements and text",
				NodeFilterShowElement | NodeFilterShowText,
				[]NodeType{ElementNode, TextNode}, // elem and text
			},
			{
				"Text and comments",
				NodeFilterShowText | NodeFilterShowComment,
				[]NodeType{TextNode, CommentNode}, // text and comment
			},
			{
				"All node types",
				NodeFilterShowAll,
				[]NodeType{ElementNode, TextNode, CommentNode}, // all
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				iterator := doc.CreateNodeIterator(root, tc.whatToShow, nil)

				foundTypes := map[NodeType]int{}
				for {
					node := iterator.NextNode()
					if node == nil {
						break
					}
					// Skip the root element itself for this test
					if node != root {
						foundTypes[node.NodeType()]++
					}
				}

				// Verify we found the expected types
				for _, expectedType := range tc.expected {
					if foundTypes[expectedType] == 0 {
						t.Errorf("Expected to find node type %d, but didn't", expectedType)
					}
				}

				// Verify we didn't find unexpected types
				for foundType := range foundTypes {
					expected := false
					for _, expectedType := range tc.expected {
						if foundType == expectedType {
							expected = true
							break
						}
					}
					if !expected {
						t.Errorf("Found unexpected node type %d", foundType)
					}
				}
			})
		}
	})
}

func BenchmarkNodeFilterOperations(b *testing.B) {
	doc := NewDocument()
	elem := doc.CreateElement("div")

	b.Run("NodeFilterFunc.AcceptNode", func(b *testing.B) {
		filter := NodeFilterFunc(func(node Node) int {
			return NodeFilterAccept
		})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = filter.AcceptNode(elem)
		}
	})

	b.Run("whatToShowFilter", func(b *testing.B) {
		filter := whatToShowFilter(NodeFilterShowElement)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = filter.AcceptNode(elem)
		}
	})

	b.Run("combinedFilter", func(b *testing.B) {
		customFilter := NodeFilterFunc(func(node Node) int {
			return NodeFilterAccept
		})
		filter := combinedFilter(NodeFilterShowElement, customFilter)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = filter.AcceptNode(elem)
		}
	})
}

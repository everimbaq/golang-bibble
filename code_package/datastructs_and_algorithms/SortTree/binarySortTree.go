package SortTree



type BinarySortTree struct {
	root   *BinarySortTreeNode
	height int
	size   int
}

const (
	InOrder  = "in"
	PreOrder = "pre"
	PostOrder = "post"
)

func NewSortTree(data int) *BinarySortTree {
	node := NewNode(data)
	return NewSortTreeWithNode(node)
}

func NewSortTreeWithNode(root *BinarySortTreeNode) *BinarySortTree {
	return &BinarySortTree{root: root}
}

// 查询一个数所在节点，返回该节点和查询次数
func (tree *BinarySortTree) Find(e int) (*BinarySortTreeNode, int) {
	if tree.root == nil {
		return nil, 0
	}
	return findEqual(e, tree.root, 0)
}
func (tree *BinarySortTree) Insert(e int) {
	node := NewNode(e)
	if tree.root == nil {
		tree.root = node
	} else {
		tree.root.InsertNode(node)
	}
}

func (tree *BinarySortTree)Order(ordertype string) *[]int {
	treelist := &([]int{})
	if tree.root == nil{
		return treelist
	}

	switch ordertype {
	case InOrder:
		inOrderNode(tree.root, treelist)
	case PreOrder:
		preOrderNode(tree.root, treelist)
	case PostOrder:
		postOrderNode(tree.root, treelist)

	}
	return treelist

}




func preOrderNode(node *BinarySortTreeNode, treelist *[]int)  {
	if node == nil{
		return
	}
	*treelist = append(*treelist, node.data)
	preOrderNode(node.GetLChild(), treelist)
	preOrderNode(node.GetRChild(), treelist)
}



func inOrderNode(node *BinarySortTreeNode, treelist *[]int)  {
	if node == nil{
		return
	}
	inOrderNode(node.GetLChild(), treelist)
	*treelist = append(*treelist, node.data)
	inOrderNode(node.GetRChild(), treelist)
}

func postOrderNode(node *BinarySortTreeNode, treelist *[]int)  {
	if node == nil{
		return
	}
	postOrderNode(node.GetLChild(), treelist)
	postOrderNode(node.GetRChild(), treelist)
	*treelist = append(*treelist, node.data)
}


func findEqual(e int, node *BinarySortTreeNode, times int) (*BinarySortTreeNode, int ){
	if e == node.GetData() {
		return node, times+1
	}else if e > node.GetData() && node.HasRChild(){
		return findEqual(e, node.GetRChild(), times +1 )
	}else if e < node.GetData() && node.HasLChild(){
		return findEqual(e, node.GetLChild(), times +1 )
	}
	return nil, times+1
}


func (node *BinarySortTreeNode) InsertNode(new *BinarySortTreeNode) {
	if node.data < new.data {
		if node.HasRChild(){
			node.GetRChild().InsertNode(new)
		}else {
			node.rChild = new
			new.parent = node
		}

	} else {
		if node.HasLChild(){
			node.GetLChild().InsertNode(new)
		}else {
			node.lChild = new
			new.parent = node
		}

	}
}

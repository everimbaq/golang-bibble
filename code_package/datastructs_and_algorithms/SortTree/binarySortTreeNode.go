package SortTree

type BinarySortTreeNode struct {
	data   int
	parent *BinarySortTreeNode
	lChild *BinarySortTreeNode
	rChild *BinarySortTreeNode
	height int //以该结点为根的子树的高度
	size   int //该结点子孙数(包括结点本身)
}

func NewNode(e int) *BinarySortTreeNode {
	return &BinarySortTreeNode{data: e, size: 1, height: 0}
}

func (node *BinarySortTreeNode) GetData() int {
	if node == nil {
		return 0
	}
	return node.data
}

func (node *BinarySortTreeNode) SetData(e int) {
	node.data = e
}

func (node *BinarySortTreeNode) HasParent() bool {
	return node.parent != nil
}

func (node *BinarySortTreeNode) GetParent() *BinarySortTreeNode {
	if !node.HasParent() {
		return nil
	}
	return node.parent
}

func (node *BinarySortTreeNode) SetParent(par *BinarySortTreeNode) {
	node.parent = par
}
func (node *BinarySortTreeNode) IsLeaf() bool {
	return !node.HasLChild() && !node.HasRChild()
}

func (node *BinarySortTreeNode) IsLChild() bool {
	return node.HasParent() && node == node.parent.lChild
}

func (node *BinarySortTreeNode) IsRChild() bool {
	return node.HasParent() && node == node.parent.rChild
}
func (this *BinarySortTreeNode) HasLChild() bool {
	return this.lChild != nil
}
func (this *BinarySortTreeNode) HasRChild() bool {
	return this.rChild != nil
}

func (node *BinarySortTreeNode) GetLChild() *BinarySortTreeNode {
	if !node.HasLChild() {
		return nil
	}
	return node.lChild
}

func (node *BinarySortTreeNode) GetRChild() *BinarySortTreeNode {
	if !node.HasRChild() {
		return nil
	}
	return node.rChild
}

func (node *BinarySortTreeNode) CutOffParent() {
	if !node.HasParent() {
		return
	}
	if node.IsLChild() {
		node.parent.lChild = nil
	} else {
		node.parent.rChild = nil
	}
	node.parent.recountSizeHeight()
	node.parent = nil
	node.recountSizeHeight()
}

//设置当前结点的右孩子,返回原右孩子
func (node *BinarySortTreeNode) SetRChild(rc *BinarySortTreeNode) *BinarySortTreeNode {
	oldRC := node.rChild
	if node.HasRChild() {
		node.rChild.CutOffParent()
	}
	if rc != nil {
		rc.CutOffParent()
		node.rChild = rc
		rc.parent = node
		rc.recountSizeHeight()

	}

	return oldRC
}

func (node *BinarySortTreeNode) SetLChild(rc *BinarySortTreeNode) *BinarySortTreeNode {
	oldLC := node.lChild
	if node.HasLChild() {
		node.lChild.CutOffParent()
	}
	if rc != nil {
		rc.CutOffParent()
		node.rChild = rc
		rc.parent = node
		rc.recountSizeHeight()

	}

	return oldLC
}

// 递归更新当前节点及其祖先节点高度
func (node *BinarySortTreeNode) recountSizeHeight() {

}



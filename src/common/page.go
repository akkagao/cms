package common

type Pager struct {
	page  int
	rows  int
	begin int
}

func NewPager(pageNum int, rowsNum int) *Pager {
	p := &Pager{page: pageNum, rows: rowsNum}
	p.begin = (pageNum - 1) * rowsNum
	return p
}
func (this *Pager) GetBegin() int {
	return this.begin
}

func (this *Pager) GetLen() int {
	return this.rows
}

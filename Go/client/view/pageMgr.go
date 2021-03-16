package view

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type PageMgr struct {
	curPage  *Page
	pages    map[string]*Page
	inputStr string
}

func NewPageMgr() *PageMgr {
	return &PageMgr{
		curPage: nil,
		pages:   make(map[string]*Page, DEFAULT_PAGE_MAP_CAP),
	}
}

func (pMgr *PageMgr) Run() {
	pMgr.curPage.Show()
	for {
		fmt.Scanln(&pMgr.inputStr)
		index, err := strconv.Atoi(pMgr.inputStr)
		if err != nil || index <= 0 {
			fmt.Println("输入有误，请重新输入：")
			pMgr.inputStr = ""
			continue
		}
		pMgr.curPage.SelectOption(index)
		pMgr.inputStr = ""
	}
}

func (pMgr *PageMgr) AddPage(tag string, head string, description string, parentTag string) (p *Page) {
	_, err := pMgr.GetPageByTag(tag)
	if err == nil {
		err = errors.New("AddPage parentTag参数错误, parentTag对应页面已存在" + parentTag)
		return
	}
	var parent *Page
	if len(parentTag) != 0 {
		parent, err = pMgr.GetPageByTag(parentTag)
		if err != nil {
			err = errors.New("AddPage parentTag参数错误, parentTag对应页面不存在" + parentTag)
			return
		}
	} else {
		parent = nil
	}

	p = NewPage(head, description, parent)
	if parent == nil {
		pMgr.curPage = p
	}
	pMgr.pages[tag] = p
	return
}

func (pMgr *PageMgr) GetPageByTag(tag string) (p *Page, err error) {
	p, ok := pMgr.pages[tag]
	if !ok {
		err = errors.New("GetPageByTag tag参数错误, tag对应页面不存在" + tag)
		return
	} else {
		err = nil
		return
	}
}

func (pMgr *PageMgr) TurnToPage(tag string) (err error) {
	p, err := pMgr.GetPageByTag(tag)
	if err != nil {
		log.Println("TurnToPage tag参数错误, tag对应页面不存在")
		return
	}
	pMgr.curPage = p
	p.Show()
	return
}

func (pMgr *PageMgr) GoBack() {
	p := pMgr.curPage.GetParent()
	pMgr.curPage = p
	pMgr.curPage.Show()
}

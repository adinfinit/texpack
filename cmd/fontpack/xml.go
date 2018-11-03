package main

import "encoding/xml"

type Font struct {
	XMLName  xml.Name `xml:"font"`
	Info     Info     `xml:"info"`
	Common   Common   `xml:"common"`
	Pages    Pages    `xml:"pages"`
	Chars    Chars    `xml:"chars"`
	Kernings Kernings `xml:"kernings"`
}

type Info struct {
	Face     string `xml:"face,attr"`
	Size     string `xml:"size,attr"`
	Bold     string `xml:"bold,attr"`
	Italic   string `xml:"italic,attr"`
	Chasrset string `xml:"chasrset,attr"`
	Unicode  string `xml:"unicode,attr"`
	StretchH string `xml:"stretchH,attr"`
	Smooth   string `xml:"smooth,attr"`
	Aa       string `xml:"aa,attr"`
	Padding  string `xml:"padding,attr"`
	Spacing  string `xml:"spacing,attr"`
}

type Common struct {
	LineHeight string `xml:"lineHeight,attr"`
	Base       string `xml:"base,attr"`
	ScaleW     string `xml:"scaleW,attr"`
	ScaleH     string `xml:"scaleH,attr"`
	Pages      string `xml:"pages,attr"`
	Packed     string `xml:"packed,attr"`

	AlphaChnl string `xml:"alphaChnl,attr"`
	RedChnl   string `xml:"redChnl,attr"`
	GreenChnl string `xml:"greenChnl,attr"`
	BlueChnl  string `xml:"blueChnl,attr"`
}

type Pages struct {
	Page []Page `xml:"page"`
}

type Page struct {
	ID   string `xml:"id,attr"`
	File string `xml:"file,attr"`
}

type Chars struct {
	Count string `xml:"count,attr"`
	Char  []Char `xml:"char"`
}

type Char struct {
	ID       string `xml:"id,attr"`
	X        string `xml:"x,attr"`
	Y        string `xml:"y,attr"`
	Width    string `xml:"width,attr"`
	Height   string `xml:"height,attr"`
	Xoffset  string `xml:"xoffset,attr"`
	Yoffset  string `xml:"yoffset,attr"`
	Xadvance string `xml:"xadvance,attr"`
	Page     string `xml:"page,attr"`
	Chnl     string `xml:"chnl,attr"`
	Letter   string `xml:"letter,attr"`
}

type Kernings struct {
	Count   string    `xml:"count,attr"`
	Kerning []Kerning `xml:"kerning"`
}

type Kerning struct {
	First  string `xml:"first,attr"`
	Second string `xml:"second,attr"`
	Amount string `xml:"amount,attr"`
}

/*
<font>
    <info face="Desyrel" size="70" bold="0" italic="0" chasrset="" unicode="0" stretchH="100" smooth="1" aa="1" padding="0,0,0,0" spacing="1,1"/>
    <common lineHeight="87" base="61" scaleW="512" scaleH="512" pages="1" packed="0"/>
    <pages>
		<page id="0" file="desyrel.png"/>
		<page id="1" file="desyrel.png"/>
    </pages>
    <chars count="95">
		<char id="102" x="1" y="1" width="38" height="74" xoffset="2" yoffset="9" xadvance="28" page="0" chnl="0" letter="f"/>
		<char id="102" x="1" y="1" width="38" height="74" xoffset="2" yoffset="9" xadvance="28" page="0" chnl="0" letter="f"/>
    </chars>
    <kernings count="1816">
		<kerning first="102" second="102" amount="2"/>
		<kerning first="102" second="102" amount="2"/>
    </kernings>
</font>
*/

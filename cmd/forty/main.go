package main

import (
	"fmt"
	"log"

	"github.com/fogleman/rush"
)

func process(number int, desc []string) {
	board, err := rush.NewBoard(desc)
	if err != nil {
		log.Fatal(err)
	}
	solution := board.Solve()
	fmt.Println(solution)
}

func main() {
	for i, desc := range Levels {
		process(i+1, desc)
	}
}

var Levels = [][]string{
	{
		"..B.CC",
		"..B...",
		"AAB...",
		"DDD..E",
		".....E",
		".....E",
	},
	{
		"..B...",
		"..B..C",
		"..BAAC",
		"DDDE.C",
		"...EFF",
		"......",
	},
	{
		"BCDDE.",
		"BCFFE.",
		"AAGH..",
		"..GHI.",
		"..JHI.",
		"..J...",
	},
	{
		"......",
		"..B.C.",
		"AAB.CD",
		"EEE.CD",
		"FGH.II",
		"FGH.JJ",
	},
	{
		"..BBB.",
		"CCC.D.",
		"AA..DE",
		"FFF.DE",
		"G.H.II",
		"G.H.JJ",
	},
	{
		"......",
		"..B..C",
		"AAB..C",
		"D.BEEF",
		"DGGHIF",
		"JJJHIF",
	},
	{
		"..BBC.",
		".DDDCE",
		"..AACE",
		"..FGGG",
		"..FHII",
		"...HJJ",
	},
	{
		"BBCDEE",
		"FFCDGH",
		".IAAGH",
		".I....",
		".I.JJJ",
		"......",
	},
	{
		"BCCD..",
		"B.ED.F",
		"AAE..F",
		"GGHHHI",
		"..J..I",
		"..J...",
	},
	{
		".BCCDD",
		".B...E",
		".B.AAE",
		".F.GHH",
		".F.GIJ",
		".KKKIJ",
	},
	{
		"BBBC.D",
		"..EC.D",
		"AAE..F",
		"..EGGF",
		"HHH.I.",
		".JJ.I.",
	},
	{
		"BBCDE.",
		"F.CDE.",
		"FAADEG",
		".....G",
		".HII.G",
		".H.JJJ",
	},
	{
		".BBCCC",
		"DDDEFG",
		"AAHEFG",
		"I.HJJK",
		"ILL..K",
		"MMNNNK",
	},
	{
		"BBC...",
		"D.CEEE",
		"D.AAF.",
		"....F.",
		"GGHHHI",
		".....I",
	},
	{
		"BCDDEE",
		"BC.FFF",
		"BG.AAH",
		"IG.JJH",
		"IKKKLM",
		"NNN.LM",
	},
	{
		"...BCC",
		"...BDD",
		"AAEFGH",
		"IIEFGH",
		"JKKKLM",
		"JNNNLM",
	},
	{
		"BBBCCD",
		"EEFGGD",
		"AAFH.I",
		"...HJI",
		"KLLHJ.",
		"KMMMNN",
	},
	{
		".BCCCD",
		".B...D",
		".AAEFG",
		"HHIEFG",
		"J.IKKG",
		"J.ILLL",
	},
	{
		"..BCCC",
		"..B..D",
		"AAE..D",
		"FFE.GH",
		".IJJGH",
		".IKKKH",
	},
	{
		"...BCC",
		"...B..",
		"AADE.F",
		"..DEGF",
		"HIIIGK",
		"H....K",
	},
	{
		"BBCD..",
		"E.CD..",
		"EAAD..",
		"E..FFF",
		"GGG.HI",
		"JJ..HI",
	},
	{
		".BCCDE",
		".BFGDE",
		"AAFGHI",
		"J.KLHI",
		"J.KL.I",
		"JMMMNN",
	},
	{
		"B.C.DD",
		"B.CEEE",
		"FAAG.H",
		"FI.G.H",
		"FIJJKL",
		"MMNNKL",
	},
	{
		"B.CDDD",
		"B.CE..",
		"BAAE..",
		"FFGG.H",
		".....H",
		"IIJJ.H",
	},
	{
		"BB.CDE",
		"...CDE",
		"FAAC.G",
		"F.HIIG",
		"JJH..G",
		"KKH...",
	},
	{
		"B..C..",
		"B..CDD",
		"BAAE..",
		"..FEGG",
		"..FHHI",
		"..F..I",
	},
	{
		"BBC.DD",
		"E.C...",
		"E.AAF.",
		"EGGGFH",
		"...IFH",
		"...IJJ",
	},
	{
		"BBCCCD",
		"E..FFD",
		"E..AAG",
		"HHIIJG",
		"KKL.JG",
		"..LMMM",
	},
	{
		"BBC.DE",
		"..C.DE",
		"..CAAE",
		"...FGG",
		"HIIF..",
		"H..F..",
	},
	{
		"B.CCC.",
		"BDDDEF",
		"AAGHEF",
		"IIGHEF",
		"...JKK",
		".LLJ..",
	},
	{
		"BCDDEF",
		"BC..EF",
		"GAAH.I",
		"GJJH.I",
		"G.KLLL",
		"MMKNNN",
	},
	{
		"BBBCDE",
		"FGGCDE",
		"F.AADE",
		"HHI...",
		".JIKK.",
		".JLLMM",
	},
	{
		"BBCCCD",
		"E.FFGD",
		"E.AAGH",
		"IIJKKH",
		"LLJM..",
		"NNNM..",
	},
	{
		"BCCD.E",
		"B.FD.E",
		"AAFD.G",
		"HHIIJG",
		"..K.J.",
		"..KLLL",
	},
	{
		"BB.C.D",
		"EF.C.D",
		"EFAAGH",
		"IJJKGH",
		"I.LKG.",
		"I.LMM.",
	},
	{
		"BCCD..",
		"B..DEE",
		"BAAF..",
		"..GFHH",
		"..GIIJ",
		"..G..J",
	},
	{
		"BBBCDE",
		"FGGCDE",
		"F.AAD.",
		"HHI...",
		".JI.KK",
		".JLLMM",
	},
	{
		"BCDDE.",
		"BCF.EG",
		"B.FAAG",
		"HHHI.G",
		"..JIKK",
		"LLJMM.",
	},
	{
		"..BCCC",
		"..BDEE",
		"FAADGH",
		"FI..GH",
		"JIKKGL",
		"JMMNNL",
	},
	{
		"BCCDDD",
		"B.EEF.",
		"AAG.F.",
		"HHGIIJ",
		".KKL.J",
		"MMML.J",
	},
}

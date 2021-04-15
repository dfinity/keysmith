package principal

import (
	"testing"
)

var validIds = [...]string{

	// Management
	"aaaaa-aa",

	// Opaque
	"rwlgt-iiaaa-aaaaa-aaaaa-cai",
	"rrkah-fqaaa-aaaaa-aaaaq-cai",
	"ryjl3-tyaaa-aaaaa-aaaba-cai",
	"r7inp-6aaaa-aaaaa-aaabq-cai",
	"rkp4c-7iaaa-aaaaa-aaaca-cai",
	"rno2w-sqaaa-aaaaa-aaacq-cai",
	"renrk-eyaaa-aaaaa-aaada-cai",
	"rdmx6-jaaaa-aaaaa-aaadq-cai",

	// Self Authenticating
	"t2kpu-6xt6l-tyb3d-rll2p-irv5c-no5nd-h6spj-jsetq-bmqdz-iap77-pqe",
	"t2kpu-6xt6l-tyb3d-rll2p-irv5c-no5nd-h6spj-jsetq-bmqdz-iap77-pqe",
	"ncpzz-qv6a2-ult3n-4mmvz-gdrhx-ileg4-upz7f-7zfqi-affpy-zlkid-hae",
	"oz7pj-bab4p-7iyos-7f3be-i565q-xib6s-pee2u-5l6wx-fexwx-ecv3e-fae",
	"xxdyl-bobgu-u3q5w-r67od-b2tc3-mkyv3-m4hlt-5psfn-6rkxt-eyv4x-kae",
	"tbl5g-hjbsu-wkhvn-djwco-uztju-cate7-g4p5b-bvqxa-5bg6a-yb6s3-wqe",
	"yksvc-462au-rmv5f-w6jlz-7h7ts-pkob5-adhrm-aj7gb-dwylz-3t53y-uae",
	"nhiqy-qh3xw-v2s6h-mkwpc-7kmo6-ue3i5-qpjk3-mqfmv-dnvox-hux72-5ae",
	"7dm5n-v2brn-p2mla-3t5cw-ppdog-ro5fo-4psh3-zotg2-ix6or-26au2-iae",

	// Anonymous
	"2vxsx-fae",
}

func TestFromString(test *testing.T) {
	for _, validId := range validIds {
		_, err := FromString(validId)
		if err != nil {
			test.Fatal(err)
		}
	}
}

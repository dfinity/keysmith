#!/usr/bin/env bats

source $BATS_SUPPORT/load.bash

load util/_
load util/assert

setup() {
    cd $(mktemp -d -t keysmith-e2e-XXXXXXXX)
    echo 'verb bottom twelve symptom plastic believe beach cargo inherit viable dice loop' > seed.txt
}

teardown() {
    rm seed.txt
}

@test "Can generate the seed phrase" {
    assert_command $keysmith generate -o seed-1.txt
    assert_command wc -l seed-1.txt
    assert_eq "1 seed-1.txt"
    assert_command wc -w seed-1.txt
    assert_eq "12 seed-1.txt"
}

@test "Cannot overwrite the seed phrase" {
    assert_command_fail $keysmith generate
    assert_eq "Error: Output file already exists: seed.txt"
}

@test "Can derive the private key" {
    assert_command $keysmith private-key
    assert_command cat identity.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIAgy7nZEcVHkQ4Z1Kdqby8SwyAiyKDQmtbEHTIM+WNeBoAcGBSuBBAAK
oUQDQgAEgO87rJ1ozzdMvJyZQ+GABDqUxGLvgnAnTlcInV3NuhuPv4O3VGzMGzeB
N3d26cRxD99TPtm8uo2OuzKhSiq6EQ==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=0 -o=identity-0.pem
    assert_command cat identity-0.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIAgy7nZEcVHkQ4Z1Kdqby8SwyAiyKDQmtbEHTIM+WNeBoAcGBSuBBAAK
oUQDQgAEgO87rJ1ozzdMvJyZQ+GABDqUxGLvgnAnTlcInV3NuhuPv4O3VGzMGzeB
N3d26cRxD99TPtm8uo2OuzKhSiq6EQ==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=1 -o=identity-1.pem
    assert_command cat identity-1.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIE8w9tDe+X5FBMP14TBA/E3gAy/N/8BiBxQR2NB0L1B7oAcGBSuBBAAK
oUQDQgAEcNM2gNSiAdbF7xf8gFK7PiNWm8DRr+hCzsFsAEFTvtZp6jypJ3f3Xhxv
o9L9hwq3YMvKasyS1vxw4slTdoRUkQ==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=2 -o=identity-2.pem
    assert_command cat identity-2.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIHDKbrXnyrCZpn8oLnf/Aly+GjkJvAayTYayMTmoLl9AoAcGBSuBBAAK
oUQDQgAE7FiHvFu/N7Fi8MRWCsLZ0Q3dcAswvwZMEzmaHLzzZu1pG11rO5NE60Tl
mKp0Tkab/fs2fifq0HmIqUrRnH5bXw==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=3 -o=identity-3.pem
    assert_command cat identity-3.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIB3B3+x3S4wvAS6yyypLROVApLNmg8AFWqgWK8roc9CooAcGBSuBBAAK
oUQDQgAE4p5mPhOCFsFFtGFIuQl1hr4JBopajLW9VbFlWMQs9y/h2QuTOyv5nS+y
mytJ2LxkaViJHzERQUskA0Ihc6GNBQ==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=4 -o=identity-4.pem
    assert_command cat identity-4.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIPHd9uF+unjyQNs5lOUVJB5K0pjwGsdi8LjR4z7Yvg+joAcGBSuBBAAK
oUQDQgAEVlxWQ+Ly7lgtMXHRSg/5UQLRgT8fio3cG/stKRXTUQZ78hM3pHYvSLND
G0skaaUZtGfppLY3mZNr1ZiG1cRX3A==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=5 -o=identity-5.pem
    assert_command cat identity-5.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIGw7GKDWzdDTtPZWMYWTxcNsCViL6ZQ8UPqGtUFY7zYwoAcGBSuBBAAK
oUQDQgAE2l9VqsOtEYPdUgNt/Zvnb4fE3Wo3FD1VMJL2rdAW591LeRRUuGQ/Kkv+
I7kZp91ZWVQPd2GrDNUXD7rSI/74Xw==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=6 -o=identity-6.pem
    assert_command cat identity-6.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIPcrApGrVb3jlhPTKMh+vqEPxvLNLxP7zZzP1yar7xoNoAcGBSuBBAAK
oUQDQgAEd/n6hcKJObZPzeD9bOAiZa1ZLEslqk3gSntGHdaLscX0fPSJKBBzrpQK
QSp1rNZDIkRJv4UWTpugn8IG246pfw==
-----END EC PRIVATE KEY-----" "$stdout"

    assert_command $keysmith private-key -i=7 -o=identity-7.pem
    assert_command cat identity-7.pem
    assert_eq "-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIHbt5vrScCNACRFPQrXgUDuRBTlyl7jpMq85DsQOUFL1oAcGBSuBBAAK
oUQDQgAExJNfFJlg1pn+7/IcDbiN8AHYZ02qBfZT+X3Vjk6s7Ztu+4LCXFjg68cJ
TiijK7kVlgFK8C24XOgK1DIXTVg7cw==
-----END EC PRIVATE KEY-----" "$stdout"
}

@test "Cannot overwrite the private key" {
    assert_command $keysmith private-key
    assert_command_fail $keysmith private-key
    assert_eq "Error: Output file already exists: identity.pem"
}

@test "Can derive the extended public key" {
    assert_command $keysmith x-public-key
    assert_eq "xpub6CNy5DbAcCUaWvFJEwNzFfQGkHJkiLT3nhsLjjBgke9twPepgyMmpAU1vKq4KnEqG6BeyoQx2YVjFuo5jSWjok4zCCNE8VDgSrZPYvGPkch"
}

@test "Can derive the public key" {
    assert_command $keysmith public-key
    assert_eq "0480ef3bac9d68cf374cbc9c9943e180043a94c462ef8270274e57089d5dcdba1b8fbf83b7546ccc1b3781377776e9c4710fdf533ed9bcba8d8ebb32a14a2aba11"
    assert_command $keysmith public-key -i=0
    assert_eq "0480ef3bac9d68cf374cbc9c9943e180043a94c462ef8270274e57089d5dcdba1b8fbf83b7546ccc1b3781377776e9c4710fdf533ed9bcba8d8ebb32a14a2aba11"
    assert_command $keysmith public-key -i=1
    assert_eq "0470d33680d4a201d6c5ef17fc8052bb3e23569bc0d1afe842cec16c004153bed669ea3ca92777f75e1c6fa3d2fd870ab760cbca6acc92d6fc70e2c95376845491"
    assert_command $keysmith public-key -i=2
    assert_eq "04ec5887bc5bbf37b162f0c4560ac2d9d10ddd700b30bf064c13399a1cbcf366ed691b5d6b3b9344eb44e598aa744e469bfdfb367e27ead07988a94ad19c7e5b5f"
    assert_command $keysmith public-key -i=3
    assert_eq "04e29e663e138216c145b46148b9097586be09068a5a8cb5bd55b16558c42cf72fe1d90b933b2bf99d2fb29b2b49d8bc646958891f3111414b2403422173a18d05"
    assert_command $keysmith public-key -i=4
    assert_eq "04565c5643e2f2ee582d3171d14a0ff95102d1813f1f8a8ddc1bfb2d2915d351067bf21337a4762f48b3431b4b2469a519b467e9a4b63799936bd59886d5c457dc"
    assert_command $keysmith public-key -i=5
    assert_eq "04da5f55aac3ad1183dd52036dfd9be76f87c4dd6a37143d553092f6add016e7dd4b791454b8643f2a4bfe23b919a7dd5959540f7761ab0cd5170fbad223fef85f"
    assert_command $keysmith public-key -i=6
    assert_eq "0477f9fa85c28939b64fcde0fd6ce02265ad592c4b25aa4de04a7b461dd68bb1c5f47cf489281073ae940a412a75acd643224449bf85164e9ba09fc206db8ea97f"
    assert_command $keysmith public-key -i=7
    assert_eq "04c4935f149960d699feeff21c0db88df001d8674daa05f653f97dd58e4eaced9b6efb82c25c58e0ebc7094e28a32bb91596014af02db85ce80ad432174d583b73"
}

@test "Can derive the address" {
    assert_command $keysmith address
    assert_eq "abde4f2523cc796bfd63564124b1c9b577c183b3"
    assert_command $keysmith address -i=0
    assert_eq "abde4f2523cc796bfd63564124b1c9b577c183b3"
    assert_command $keysmith address -i=1
    assert_eq "0b98ad668da5702f2d127c8de01cbd0de3ed5ce8"
    assert_command $keysmith address -i=2
    assert_eq "a5ebdb7665b68446328e8356c1808db3fc068d93"
    assert_command $keysmith address -i=3
    assert_eq "50a7bc9a4d205a8ce96e6c1bd3a6543f6acb19b4"
    assert_command $keysmith address -i=4
    assert_eq "985f1a72ebf6580770d9047d78823e8be5777a39"
    assert_command $keysmith address -i=5
    assert_eq "6c22e8745c81b5cffcc0b12383d3668623c40172"
    assert_command $keysmith address -i=6
    assert_eq "465111daace75b765d2df70619787e0f5bec2e82"
    assert_command $keysmith address -i=7
    assert_eq "db6de8140a50b31ed18494e93b03cfe5ba082d23"
}

@test "Can derive the principal" {
    assert_command $keysmith principal
    assert_eq "t2kpu-6xt6l-tyb3d-rll2p-irv5c-no5nd-h6spj-jsetq-bmqdz-iap77-pqe"
    assert_command $keysmith principal -i=0
    assert_eq "t2kpu-6xt6l-tyb3d-rll2p-irv5c-no5nd-h6spj-jsetq-bmqdz-iap77-pqe"
    assert_command $keysmith principal -i=1
    assert_eq "ncpzz-qv6a2-ult3n-4mmvz-gdrhx-ileg4-upz7f-7zfqi-affpy-zlkid-hae"
    assert_command $keysmith principal -i=2
    assert_eq "oz7pj-bab4p-7iyos-7f3be-i565q-xib6s-pee2u-5l6wx-fexwx-ecv3e-fae"
    assert_command $keysmith principal -i=3
    assert_eq "xxdyl-bobgu-u3q5w-r67od-b2tc3-mkyv3-m4hlt-5psfn-6rkxt-eyv4x-kae"
    assert_command $keysmith principal -i=4
    assert_eq "tbl5g-hjbsu-wkhvn-djwco-uztju-cate7-g4p5b-bvqxa-5bg6a-yb6s3-wqe"
    assert_command $keysmith principal -i=5
    assert_eq "yksvc-462au-rmv5f-w6jlz-7h7ts-pkob5-adhrm-aj7gb-dwylz-3t53y-uae"
    assert_command $keysmith principal -i=6
    assert_eq "nhiqy-qh3xw-v2s6h-mkwpc-7kmo6-ue3i5-qpjk3-mqfmv-dnvox-hux72-5ae"
    assert_command $keysmith principal -i=7
    assert_eq "7dm5n-v2brn-p2mla-3t5cw-ppdog-ro5fo-4psh3-zotg2-ix6or-26au2-iae"
}

@test "Can derive the account" {
    assert_command $keysmith account
    assert_eq "53a3ef3b11b69c6411cd1667970e24577c857acf3cdd67436656c75d4ddd3cc7"
    assert_command $keysmith account -i=0
    assert_eq "53a3ef3b11b69c6411cd1667970e24577c857acf3cdd67436656c75d4ddd3cc7"
    assert_command $keysmith account -i=1
    assert_eq "6bf62b34a4eeb32721aeb99fe598d760ca252a0324155fc8b0ef203f3fc620cc"
    assert_command $keysmith account -i=2
    assert_eq "2f1a51eadb26c0dcee782a8069018abdcce2c60a386ad9c0351ac4328d7346d8"
    assert_command $keysmith account -i=3
    assert_eq "927e01c4b46c7f4d3363c9823de9af067eb0f3e13e7b4cdb455bedbe7268bf0e"
    assert_command $keysmith account -i=4
    assert_eq "2b84c97f17ea887eb0374d14db1402867879a163c4799583327bd38afeb93d51"
    assert_command $keysmith account -i=5
    assert_eq "e29bf07177868d4c43b08857eb47dd99991aafafa918290e7a2561deab054aa1"
    assert_command $keysmith account -i=6
    assert_eq "9fed555b071f02563c0b24a3a70f5b2a1a863b97b80717da70e4816d1edc47e0"
    assert_command $keysmith account -i=7
    assert_eq "91e6f4b6bb528d8a26288cbc677407e09fb19493a567aaae7589ac0b1d03f9dc"
}

# keysmith network service API document
this is a project that forks to [dfinity keysmith](https://github.com/dfinity/keysmith),Basically use GIN to provide a network service.

**Account**
-----
  Return your account identifier.

* **URL**

  /account

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  * **mnemonic** `{ mnemonic: april hello table melt soccer fork capital valid two loud govern hammer }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ code:200, msg:"", data:{accountId: "4cdfebb72b356228e4893ac42e62af8d4c9b994044294558ce678220fe4a0f07"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:Invalid mnenomic, data:{}" }`

**Generate**
-----
  Generate your mnemonic seed.

* **URL**

   /generate

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"code":200,"msg":"success","data":{"mnemonic":"tray eagle canvas belt three diary rebuild result polar situate hockey blood"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:"error message" data:{}" }`
    
**Legacy-address**
-----
  Return your legacy address.

* **URL**

  /legacyAddress

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  * **mnemonic** `{ mnemonic: april hello table melt soccer fork capital valid two loud govern hammer }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ code:200, msg:"", data:{legacyAddress: "26d5c0383c71a053792d6cdfb677e1f30ee887d2"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:Invalid mnenomic, data:{}" }`
    **Account**

**Principal**
-----
  Return your principal identifier.

* **URL**

  /principal

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  * **mnemonic** `{ mnemonic: april hello table melt soccer fork capital valid two loud govern hammer }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ code:200, msg:"", data:{principal: "2f442-hk6ph-lhi2f-ocdcz-welrk-2qyqb-byk3e-35jud-zkpj3-uf4i2-pae"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:Invalid mnenomic, data:{}" }`

**Private Key**
-----
  return your your private key.

* **URL**

  /privateKey

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  * **mnemonic** `{ mnemonic: april hello table melt soccer fork capital valid two loud govern hammer }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ code:200, msg:"", data:{privateKey: "-----BEGIN EC PARAMETERS-----\nBgUrgQQACg==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIAd1oPF3WdRcGVB5ONAjA6jiZrZyqMeSrEq1RyEnWRcooAcGBSuBBAAK\noUQDQgAERz4oHpj+22gBKceoTUv8mHoa+ufI6ntSlFmKdOv82C2z2n7CiUtTbRaZ\nd5PuDIAACaah7/CbjXObbng0EEEPrw==\n-----END EC PRIVATE KEY-----\n"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:Invalid mnenomic, data:{}" }`
    
**Public Key**
-----
  Return your public key.

* **URL**

  /publicKey

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  * **mnemonic** `{ mnemonic: april hello table melt soccer fork capital valid two loud govern hammer }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ code:200, msg:"", data:{publicKey: "04473e281e98fedb680129c7a84d4bfc987a1afae7c8ea7b5294598a74ebfcd82db3da7ec2894b536d16997793ee0c800009a6a1eff09b8d739b6e783410410faf"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:Invalid mnenomic, data:{}" }`
    
**X Public Key**
-----
  Return your extended public key.

* **URL**

  /xPublicKey

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  * **mnemonic** `{ mnemonic: april hello table melt soccer fork capital valid two loud govern hammer }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ code:200, msg:"", data:{accountId: "xpub6BqaNJo8GBEPGoaSkVBSjEdgUKLuTjUFB7wo2u4vbd5bPo7BNAaGmh1BqDBQ7u9gaQ63XGrRjz3dgcukifNerBK54LtLhKcFTRPjWNvR4mN"}}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ code:400, msg:Invalid mnenomic, data:{}" }`
package bls

var Scale2RootOfUnity []Big

var ZERO, ONE, TWO Big
var MODULUS_MINUS1, MODULUS_MINUS1_DIV2, MODULUS_MINUS2 Big
var INVERSE_TWO Big

func ToBigNum(v string) (out Big) {
	BigNum(&out, v)
	return
}

func initGlobals() {

	// MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513
	// PRIMITIVE_ROOT = 5
	// [pow(PRIMITIVE_ROOT, (MODULUS - 1) // (2**i), MODULUS) for i in range(32)]
	Scale2RootOfUnity = []Big{
		/* k=0          r=1          */ ToBigNum("1"),
		/* k=1          r=2          */ ToBigNum("52435875175126190479447740508185965837690552500527637822603658699938581184512"),
		/* k=2          r=4          */ ToBigNum("3465144826073652318776269530687742778270252468765361963008"),
		/* k=3          r=8          */ ToBigNum("28761180743467419819834788392525162889723178799021384024940474588120723734663"),
		/* k=4          r=16         */ ToBigNum("35811073542294463015946892559272836998938171743018714161809767624935956676211"),
		/* k=5          r=32         */ ToBigNum("32311457133713125762627935188100354218453688428796477340173861531654182464166"),
		/* k=6          r=64         */ ToBigNum("6460039226971164073848821215333189185736442942708452192605981749202491651199"),
		/* k=7          r=128        */ ToBigNum("3535074550574477753284711575859241084625659976293648650204577841347885064712"),
		/* k=8          r=256        */ ToBigNum("21071158244812412064791010377580296085971058123779034548857891862303448703672"),
		/* k=9          r=512        */ ToBigNum("12531186154666751577774347439625638674013361494693625348921624593362229945844"),
		/* k=10         r=1024       */ ToBigNum("21328829733576761151404230261968752855781179864716879432436835449516750606329"),
		/* k=11         r=2048       */ ToBigNum("30450688096165933124094588052280452792793350252342406284806180166247113753719"),
		/* k=12         r=4096       */ ToBigNum("7712148129911606624315688729500842900222944762233088101895611600385646063109"),
		/* k=13         r=8192       */ ToBigNum("4862464726302065505506688039068558711848980475932963135959468859464391638674"),
		/* k=14         r=16384      */ ToBigNum("36362449573598723777784795308133589731870287401357111047147227126550012376068"),
		/* k=15         r=32768      */ ToBigNum("30195699792882346185164345110260439085017223719129789169349923251189180189908"),
		/* k=16         r=65536      */ ToBigNum("46605497109352149548364111935960392432509601054990529243781317021485154656122"),
		/* k=17         r=131072     */ ToBigNum("2655041105015028463885489289298747241391034429256407017976816639065944350782"),
		/* k=18         r=262144     */ ToBigNum("42951892408294048319804799042074961265671975460177021439280319919049700054024"),
		/* k=19         r=524288     */ ToBigNum("26418991338149459552592774439099778547711964145195139895155358980955972635668"),
		/* k=20         r=1048576    */ ToBigNum("23615957371642610195417524132420957372617874794160903688435201581369949179370"),
		/* k=21         r=2097152    */ ToBigNum("50175287592170768174834711592572954584642344504509533259061679462536255873767"),
		/* k=22         r=4194304    */ ToBigNum("1664636601308506509114953536181560970565082534259883289958489163769791010513"),
		/* k=23         r=8388608    */ ToBigNum("36760611456605667464829527713580332378026420759024973496498144810075444759800"),
		/* k=24         r=16777216   */ ToBigNum("13205172441828670567663721566567600707419662718089030114959677511969243860524"),
		/* k=25         r=33554432   */ ToBigNum("10335750295308996628517187959952958185340736185617535179904464397821611796715"),
		/* k=26         r=67108864   */ ToBigNum("51191008403851428225654722580004101559877486754971092640244441973868858562750"),
		/* k=27         r=134217728  */ ToBigNum("24000695595003793337811426892222725080715952703482855734008731462871475089715"),
		/* k=28         r=268435456  */ ToBigNum("18727201054581607001749469507512963489976863652151448843860599973148080906836"),
		/* k=29         r=536870912  */ ToBigNum("50819341139666003587274541409207395600071402220052213520254526953892511091577"),
		/* k=30         r=1073741824 */ ToBigNum("3811138593988695298394477416060533432572377403639180677141944665584601642504"),
		/* k=31         r=2147483648 */ ToBigNum("43599901455287962219281063402626541872197057165786841304067502694013639882090"),
	}

	AsBig(&ZERO, 0)
	AsBig(&ONE, 1)
	AsBig(&TWO, 2)

	SubModBig(&MODULUS_MINUS1, &ZERO, &ONE)
	DivModBig(&MODULUS_MINUS1_DIV2, &MODULUS_MINUS1, &TWO)
	SubModBig(&MODULUS_MINUS2, &ZERO, &TWO)
	InvModBig(&INVERSE_TWO, &TWO)
}

func IsPowerOfTwo(v uint64) bool {
	return v&(v-1) == 0
}

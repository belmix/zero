package main

import (
	"encoding/hex"
	"testing"

	"github.com/dogestreet/zero/stratum"
)

func TestValidateBlocks(t *testing.T) {
	n := 200
	k := 9

	table := []struct {
		headerNonce string
		solution    string
		nbits       uint32
	}{
		{
			"04000000a2866b73bda67ab39e80e5843e169af2d7dda5c19e1117e018edace87a13f4019d745e380766d475fb4c08cbdae7b815099db423e28e5785436b2e8877e50d720000000000000000000000000000000000000000000000000000000000000000a6320c58ffff03200000000000000000000000000000000100000000000000000000000000000000",
			"007daabc87bca419f7192356a1f0ae1e451134a6b911c69cb7bec528b15a3bd1ee39390a0d94d4d8080b23c3cbb2d80f4b31590135124db32e365042fb96ac2876d66e770ec7f773ce350ce640a5b60306b11c99055f057ec55aa9e59a14d0c9249862b0e5c16d31ec0e78177b040657dd3243b5f83f5744b58b77cc8bc91bd27425cec7b940916a250529f60edb33fbdd4141341d26f2a3aff8fdae52a55b203a233deb8f1ceded0bb6e2ddbe16ddbd379712825d7a7544b9547b7950405acfecc850bc61083e368578f75455f9879ac14213ff11f1ac0fd34497e613c3c66d9039882a132f722d12655c616a0725cf0c1345b5c88da5fa02ba09701c37737243de2379a24e25fe74bf9ba6459c179c813291a436434e3e3d5b93c3c383fbf0e2a293fec3db2a240efdd69b20ab4a4133e45faf3d6e7467d95241539357bff0250143b3ef07a021f1ef4a0e3bfacfb3036156458915f09764c471ad779bc7a8862b9b3b310cff5250c2a32bfdd2534559fee3c4b181efbb88fe053545255dce838ec314041c46256a013a46ad27af2eb20c4c4d197cd9eed0c49b79d8112a5af875f0f5129e777302fa810dec352480c4ef3cf96177abb52714ebb90007a72543c8c931bbb079bad5ef7712f3a813da9ce4298c95536b94a3de747f6ca91b03b1dc9d419ee2fa9a2e063dff69a5400be32ed224cbfd630e080d2f766f44fb2b214bd3a3336bd87dd6e0fc482f2dde9ad6b3cd9c0a7cb5a33f92481ffa0f80944b311187294d29891ce5d072e2cdd02143c90bb37dadbf2d5d9710794be6f5246a835d2d30db251edcff640d0a3080b0384f6d2dbe56666365d5b97dc564bae6a24bb05504eab62d71eb2ec72c287714e77517ff941b0f3f35019e62d6f77d65933546d518277922dd37086978d45ecd2a06579ab08756ef71fb8a263512b521028e570eb94d4f95813604d6736f5e2ecea1fe8053422bcc461adeb2ef17d655d7de6afd32d4547bac85039c0daf799924acd1b5805bed46b95a0439b6b68c19b065f03551d6ad49b566636a3393bb30145e43cd0360f3045c8b3b40a26e186a4f65143e56e01d242006b6339815362c97ed7f514f927120951ccfd69820065dd339d5fb0e71e278862c347f084e39d59914fa0668e10dfe910e210e297310e3ecde52e41ed9477603187748ca8422da3c3429c837705c7f890a7e0c4a2f8784cd24e1d0f176b3230c22dccb1f2b8b5d57c225c57aa0c2138c5dc7bbd442a8b4658e32f0d929b426c22dc8292007550049e582053909db3975bb37b504a8eb90848902714d2252c95321708f47ba5f9b7e21c48fe6d95044adc27e241d6eccc4e1280c2cfdeb17ad7c04dbdade352213b758caf8d1c3968bdebf0964dd77e8c27480dbfdaa182f51e6659b61351bb88b0896360e15c50fe44d56816c3deef0f89b3e92186c096234229c82a6fa342c61b1275889a97bb7dc50700cb11218a10c80f77b5933fc5d5de7b5728bf96e1110af071cdb2372b166f8344df6c0ddf5abedfdfe040abb644575e2f915d72412eacbab885e4943bc0a580c38440674750dc1ab73e122aaaa723af1f83a8d9f1a54cb89de8e1edaba9a93e62c3f9207b8209ef4891de3174dbe952e4104d3545c18c925c226ce13701c11530a48512d02b9707d9875e36def97768d9f1fc83604abb4590ff9eb600c85d6b43925aa30321c4597181b4e41749d30471cbda6fe85f3f82ab6bc19a34e37c4cba27229a59dd14554af5b2fa615b4ee1609621542d321adefbe4f98d3230138a8e970d495bd833c15a8ab0db7888ad17d7565b5783148063acb5cc2f825b21504e3701f81efce0bd5e2e7ee3f0433f88d2a435e2a5df799832fcdc33486f8d58f6157d91c1ced",
			0x2003ffff,
		},

		{
			"04000000ea8dac2270df0ba7a032cb34d3ed54b8f6347c1ddd55b6bbd915814afcbfaf02d3304f50752e6c74019e09890dd588d028daa81d9790d4b99acf76a53937a1100000000000000000000000000000000000000000000000000000000000000000a8320c58ffff03200000000000000000000000000000000102000000000000000000000000000006",
			"0095b6ac79cf6600e8c5f5d04368177f60f69eae161aac122265647ac36ff0e56d9fb6ee2d5b7fbbc77808d16da77e0ba8b6c248932b4246d8f976199361e92c811b0b01eab8ab68daf49452cb79c30b1e5e5f7d0275dec16249cc17cb66428e8e9909ae6fff7f86a62d219c0170567759fbd5faa5676eaa4f1142db09b709842fca87dbe4374e1df378a37e88827f475fd7fb3dc9ac0782158856ca64a53f4140969dcc309c6fbf01efecb8a2d9d37d16837348e9eee3157b13fcd2a02051133ad1c89a72b5403279d6dbb574dda997d6be086c695ca64e7693cec8739200ff4ffee744395f222e46dd82552447a1ae4e9339061a52dbb895dfe88b02254bfe762393193a79b67ed0d1b35eb460f622ca30962daf2eee47a78ae5c5408539c97988222e342b0c182e911f432b9a3c1521b182e657788d6c2e27a41735d8d130e76375857b25e0bf333e3dae585ae475012c8942780052a896b1667c47ef34eab6521d53b82c16253d081d2d4b3befa5bde9391fde711cd4441809e95afd4b68d41b6f5b11e37031ff31e8ccd331220dfc6611cbdf9c21f370c281e45dacd6b7aedc0fd5088dbb659152cfeedeeaa1d2de950dc1604f4e37e420af7d9f7193188eeab5e2f7d5699acf26af3f450320b79106b657380ba281f5c81e397c02bee756979524c9c47060136608f699c660d9562d2a538a76cbb80b5d94bbe3cdae5d7e8dc4eccd534592bf50dae453201ea592a9549688af0c72a5887dc252554d1f32722338c5451951de2bc927f68efe4adad6cf2a36a3a82bf62fc6e816c44fb25bc93f4765dd7733bf1d5aa510df5d77bfccc47f0a674397f9b13c81ad9ebf989013f3b98e73354079b67442607f3d9a5d551f10a3c116a61c9bf74b6177d25014745d366819a622d069832893434ea60a9b4b110c4300fef3e34d5ce8df2f7401a5929e571e88837cb213ace8befdff2396fef0a022c65cadd4ce3fcccd5199f98bf974a2ce4079b0bf384fbf74c8946330f564b5259ab1f4ae3479bcd8fb4d2172a3c5e103a5f49a0573b0c37d3dd335b383980628067f9ca205a766f31152a5b299b10e209aa83612dfd93869182961f512b2c782d35f21cc93b314a828db52314214e371172cc476eaed95793ab7cee7d9362f949dfc4f08e1a6b3c529aac58fb1f4e21d7b3603278ab9cb56d78335356151bbaa2f84dcb29f850d141458c413f34869adc832c4c22f3e3633d9b838962964aa7de2dda419323c237e0cd60394f8367acfe93318bb667519187d325317bd894665d324d1fc93cc05c8c2b8f8e7202fafd221ac3aca0da333935a6e3c0a949307dfd48ae71d4947ab2d516232044d3aba0b0c5737b5e9a35c2bf00fd7dd6ce094eab820fb4b501d9d4c1df30d66c68a9d33d15ec69ab5e1200fe28902ee6c4673d573e552197253c9e6b75662eed3d88d1c37315db70c0b07bf9cd2f7646d1be24f51938d581864bf0cf94b959fe8f69824816446371c117fd8c827fcec2e909b3ce73fa65575c03faba25ebd3c4bd104c16ae143c60817f911529208eb670ec8b8fb1afa40456c9f59b6b537caa0c512acf7d1420fb15837680607e206e69fea11174361e3a491c75193523df9330fdf34848cc870586c6ac502cb7da55557744d61db0de2222a27c44d6dbc970584dfc2d92964068cad60310fea00181111c1090e0454123411e603dfbb7226138d46659b6ccd8b883f3185eabd33ceca3b37ecf617f79afc1dc6c9d9001e8488c2c2272d8c9b58c8ba108a6a66632a1951c1959242737e2bad56048f771d559533282e9f85494f33e5766dc0d0fe69419b83ec174c5d0a13284dab4c8974495fcb111aec0ffce1792d6fdbcfb5160e62e904d3936be251550e934abd91",
			0x2003ffff,
		},

		{
			"04000000b82169ca65fa5ca9c7cebd1aec9c9bbd7cb87304fea41916f29afdaf172d950032b3ea875acffa15e0cb806130a9dd4d174f4774e75b569e814b150bcde1ea020000000000000000000000000000000000000000000000000000000000000000ad320c58ffff03200000000000000000000000000000000101000000000000000000000000000000",
			"0035eec8f8310fa196cf3321d85a2649afacf024fe093b335fddcac7453ca820c25b759dd57a9c5df3e601688710e616bb50bf14a21e20c324f5ee49bc080074799e22cd65db89f17e48106645398642373b069c06b34317e66eec8b859ad1d4f7e169e55a9db3620b16796e6bf506fc20c48123d1d670fb91b30237438c0b09ff84f858e9b1de3ea47bb1d2f07b1294791a4f2682f97dada050137d1074bdf9f9c85da802b5d4fd0613c8c10a88a81ba8c580fb07e4d0093c50ca0814185b3e0180330b29ecefa302a13794a6b8e636775e1676efb2691af54f164ef20770d3eac67abcbe420447aeb52df4e9eb59c1b0748e80d7fd9f00a4be909d0b2952fe49c4d043188101aaa56031acd82ab55a940c58e8a8ac22eb15ea43e48cbd265e9d66651b157120d509226ea24513f35c14d3807058594c129e937344aac58c2769f32de3be050dab55f78a98e1dce92500959e9b07005838f2db112212ae8cc8f7649c649c1c8c7af1f168abc79e43380ccb4e8c161278b8de9f00a1523544eeab3db0ae901db372bc081175e71f7502c4526e2d1a428aeb3de1d0f4a9fca8d7f4b73bc715419b787587e4a4e9de9a5ad0e7b74fb64d1db6c030ecefa6a39b955d1c087481ce3496bf4721fb7ccb16a1660cdfdb9cc3ca90c90369ef0ae6b6a71f42fd1daf77f0d30a084d4c4ac71adfc888db9698fdc9460a4ee05a2261984f953b922a56ec15f0fc0aced3ec3a2103cb066fd2a5de27747b2051e9d639e6d95ea62814431bcb544b813fb2f384e7b8c964ee6c1dae443c38775d3e99814aec13d5af19c379b5a3bc9da2fe128d5b031448166309af773ac53af729ff365f05ed12f8a113372df67dbd6901a173bd22d4de20c79fe3243a4611751215db65bfc3ef3dd7d96a5c0e33c67d845bb553f7a2f0d173da9896a7c52aee356b96df3d0135f7d02bcda62fe84cc01a246dce8b0fb439fb32078d687515d17da17bd14130931048475b337d15a008689b9d99872c810477e0c0502af5de13aa54e04e379823579a9271dd09a873dfd93fafcab25c3945e504fd194cf9eb548b6735222a0b55b717a3affe5ceb14e3ea8a12cf9c03faf951faaa22d86190fe9f48c40d8a04bab0f0c21bd7e2e29855bf1b210812acae6d4fbde6bbd964fdd197563722da611265e6646fb590092e76f50a58ddc92f41763ec36920bdca90fccd4b0deddbef5aca275f3511563cdf7b9651e5b157c7ae1e23a548b8593aa3b068d39def48b811b22c9dd9b532f8f237c10d6754b9a7b564f45bc42e31a4bee3460c429377015111b4c89144a432ef64a932805b0dca455c2f54a413fb83130a55bf9757f5aeb95ab7c14411cf6ee807a5db11909c51470871cb949a197eb76818df6de9f44c337fa47a76b9edd7d6724ad47bb1ad02c7e355a28459ecb59b33b5591eda88ef04f4465911a877e8f646d44d1e238299ebb9493663861661960e835f442acd9800b7f1884137ea9eeaa5b33ac48e225c3e10241e7761e1a613bf48f7169d79069e1dd807fdffdfa6641b658a8ce523f273c95f65ca1e1401300bf59f41e05c11044605b887e2586e1c6cb27e0511eb5d56419d6e5dbf7bdb44e8e0efcf2dbd79efe2421b165ef6e35afd3de636bcbffd220dd330989f3b0cbc090f9c46f59ca2c2a62cb1b6a029dd4c8f74f51f7024ecc2de81cd9485f9efa1eeb5f33284be9450152e257af62959c568326423352ef8b9559973391839f47b343feb075ffccc73b1e0618720ee2f9af6f20dd73fbe2e8dea508e0445943e3bef15f0e99471db124e847cb0efb189d80e68fbb86c5772626935511c2c7c06f4dc54ea87343806929e42d26f56d83cbfee67f3e34997e669afe74ec8657aedb6aabc2add641e",
			0x2003ffff,
		},

		{
			"04000000615e22814c6676bf27fb3878e3f8a2291dadba9b00f08d6ed04d33dff80000002d64c2627f5f830a48806f6dcbe0ceae24b7cf16eaec0b86116ec705827799e800000000000000000000000000000000000000000000000000000000000000003c3e1158a6dd051ec6b2ff7f00000000000000007fffb2c5314c0000000000000000000000000001",
			"003339124689ffb5bbfbb1d650fd5810da85ab200c1759adc085d7e303c1a928d965665cae78dd34cba10e18f21e5889fa669bb330f27c9362e34eea3d35d711bca9edf54a63d96825c2e95c767acd92e19726ed108dc6182ac57937487c91f8ea943ebb6e9a3e1b002adb4da7aab13923fdd5785632758e4a9a60d6ce1924cff40d8c6a0a0196c0740a2226443569ed4d56c15761c4cf0edeb215cbd8462bfdd9beba5f0afb2dd70486f5f86eac1509a3f1d3d2456ee95ebf447cd851103edd86b3da63bf28d3c7c7f05606da21d37132920f943e66ea52dd78e449d1911da1d26522ec52599c1f29f36bf899fa25bca9c914b8d4159edcfdfb38490be0e8984f927adea802b0df2216e14097046aea7d14514a114d23be41b676419070f7b2ae9e65feae70181d78d063c650bb754206db1ec36eb60c51b57b4c267779993acdc23e9390e5838674a3125e473f79a104f89680a98dad41607f33fe66d417d92f4b8e5ec837c645d3a0508531b81063cf6c50ad22c8da9fd618087fcf7752ccbe0f5627c34fdd437b1e992ad97e2d227b25fd37191a58f41cd68437564e6e182f594f8c09ba4f184e1f4e719fa988014ad7ca8b078059ad4e10ce6f92f1eed5c3a2dab8fc7a716f62fb755bdba90d547aba542d0b678500c4275a7683f71f5f7dafbd8f6034ea2f2724cd5159ca8dbe7a16af2fbf9da3ba1514d674a60885c2d643e321972d962192ac54da0719ba579123923858963881ccf9789896a30a9e9b671bca4480cce977afd9f114eb6561b65281337fe90b24ef871756b0bab5f73652eff326d35925bd4f56c0157c258e320e331f209d2781817506567d1f1d9a102f278363f2f84eb3f20a53afca315319962911915f16dd92faa20b6d1e6f9ee236f1c62fc4fe95702f1526645282ebdc350988639430e5d25c75a7baf5c10b00d3faee1f4cab83f99905817aaf8c7d621150cba95b1514e680ebd98fd35517f7f447ec32d22617664c0c805d236e5840aace09f13051d717fe92ecfef4351650a16a675a5d27886a95bec6eba9a2fe609aee5f10e010c161490115e1a4b28cfcf7ed15f40f9acb3319376bc85dc9dc9ff70f63c9daab8a5262e93acb1429059d5d781bcdb9c09e43aa43400b8dc7b40eac0e2cfaaf8abbe85bcb964cf3ec1fa8e26d3eba52ac67074192088848de4717ba8480cde0a0b5abf1117b99194f0ef2433e3e7bfad775f0dc6a91df4d839d0e371bcd0a0d0c508a06d532f49ba7f187aa3adf773ee5269293150f89ee8be7723328683b57615fed5c1e5616cb7bdaeaa7e30df97dcd55b77f0b7fe1eddfbb5019b97ad3388c7095e51a51c683270f3e46ef3ab6e82590d2c6e01b317d99ce2a7ceae0a21acaa05e9df54ef4acdc7a6dd2e3e14149f00ae40f4b538cbc7c14012d18679488ec9f500b62aa594c73cdc73d1ae21c1453b3f58268aeab55431201a518b79dfc0674840b049b92b4af430ac0a787b09746ce34d683455a388e0e94970e03186993c30cd30774b8cef584618e5ae51fe29b2f48eaf6f7d71f59f41cef48473c72bf1a7a35ffbb790da249c5275ce4663bd2140a0c5257979541066355ceab5f19a99296fca9699dc3e74a5fb0fd42b67f01021407b5b64b7696da5fdd39fd48d92718058ddd13c10cb3430edf45e2ee3ef2e20a4e520b9a40def754ce54c45b20ab4dde33fbe287a746dec5f712c872fbedac2677bf1741a90ac6aa8a0fe9be766f37a78b83e69b381140708620d0734709f62a31590e068eaffad8970c015a4b8aebf66f274adb8a1c22b81276e1dff836d6b5fcfe226ea24c4c1222649a769e12885e8601ed34b5e1a625014a467e919ff4b3f5df689c1b9aab9d5f84f1eedbd83179d3e32c47fbbf66",
			0x1e05dda6,
		},
	}

	for _, b := range table {
		headerNonce, err := hex.DecodeString(b.headerNonce)
		if err != nil {
			t.Fatal(err)
		}

		solution, err := hex.DecodeString(b.solution)
		if err != nil {
			t.Fatal(err)
		}

		// Extract the target from the header
		target, err := CompactToTarget(b.nbits)
		if err != nil {
			t.Fatal(err)
		}
		res := Validate(n, k, headerNonce, solution, target, target)
		if res != ShareBlock {
			t.Fatal("Failed: got:", res)
		}
	}
}

func TestTargetCompare(t *testing.T) {
	table := []struct {
		nbits uint32
		hash  string
		ok    bool
	}{
		{
			0x1e079519,
			"00000384e920ba1255421924060add921ced0931043dc43ea1ad50e8627f118e",
			true,
		},
		{
			0x1e080508,
			"00000357bd8f61f4b82552e4b52921a77d857e300f3b83b961412dd87532dfed",
			true,
		},
		{
			0x1e1730da,
			"000012f0d19e93070d1f53db4d0a275941318647d50d86eb89c2df9e195e954f",
			true,
		},
		{
			0x1e1e240a,
			"0000184fee3f821b27af1224688aa8a8e90e36e80ea800c3e6835d28c80e2ce6",
			true,
		},
		{
			0x1d1e240a,
			"0000184fee3f821b27af1224688aa8a8e90e36e80ea800c3e6835d28c80e2ce6",
			false,
		},
	}

	for _, x := range table {
		target, err := CompactToTarget(x.nbits)
		if err != nil {
			t.Fatal(err)
		}

		value, err := stratum.HexToUint256(x.hash)
		if err != nil {
			t.Fatal(err)
		}

		if TargetCompare(value, target) > 1 {
			// Did not meet target
			if !x.ok {
				continue
			}

			t.Fatal("Failed")
		}
	}
}
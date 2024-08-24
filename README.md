# Mini Aplicação de Linha de Comando

Esta aplicação foi desenvolvida para ser executada diretamente pela linha de comando, oferecendo uma forma dinâmica e interativa de realizar consultas sobre endereços web. Com duas funcionalidades, a aplicação permite que você obtenha o IP e o nome dos servidores associados a um determinado endereço.

## Funcionalidades

A aplicação possui duas ações, ambas exigindo um endereço web como entrada:

1. **Obter IPs do Endereço Web:**

   Ao fornecer um endereço web, a aplicação retorna os IPs associados a esse endereço. Para executar essa ação, utilize o seguinte comando:

   ```bash
   .\command-line.exe ips --host amazon.com.br
   ```

   **Exemplo de retorno**
   ```cmd
   3.140.149.207
   3.143.158.216
   3.19.4.205
   ```

2. **Obter Nomes dos Servidores:**

   Ao fornecer um endereço web, a aplicação retorna os nomes dos servidores em que o endereço está hospedado. Para executar essa ação, utilize o seguinte comando:

   ```bash
   .\command-line.exe servidores --host amazon.com.br
   ```

   **Exemplo de retorno**
   ```cmd
   nsgbr.comlaude.co.uk.
   nsusa.comlaude.net.
   nssui.comlaude.ch.
   ```

## Tecnologias Utilizadas

A aplicação faz uso do pacote externo [github.com/urfave/cli](https://github.com/urfave/cli) para facilitar o gerenciamento de comandos e opções na linha de comando.

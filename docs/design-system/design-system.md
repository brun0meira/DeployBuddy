# Design System

## Sumário

1. [Contextualização](#1-contextualização)
    - 1.1 [Proposta do Design System](#11-proposta-do-design-system)
    - 1.2 [Como deve ser utilizado](#12-como-deve-ser-utilizado)
    - 1.3 [Referências](#13-referências)
2. [Validação do Design System](#2-validacao-do-design-system)
3. [Guia de estilo](#3-guia-de-estilo)
4. [Biblioteca de padroes](#4-biblioteca-de-padrões)
5. [Biblioteca de componentes](#5-biblioteca-de-componentes)
6. [Prototipacao](#6-prototipação)

## 1. Contextualização
Bem-vindo ao nosso Design System, um documento dinâmico utilizado para o desenvolvimento da interface de interação do usuário com o nossos sistema, detalhando os elementos de design que visam proporcionar uma experiência de usuário ótima e alinhada com as melhores práticas e tendências atuais.

Este guia foi criado para ajudar a manter a consistência visual na nossa solução - DeployBuddy. Aqui, você encontrará 3 principais artefatos: "Guia de estilos", "Biblioteca de componentes" e "Biblioteca de padrões".
## 1.1 Proposta do Design System
O design system proposto foi desenvolvido com o intuito de atender a necessidades tanto de designers, como de desenvolvedores, facilitando o desenvolvimento da solução DeployBuddy, bem como uma manutenção da plataforma, disponilizando insumo para:
- Design de novas features;
- Possivel aumento de escopo, tortando o design escalável;
- Manutenção futura do design;
- Facilitação do desenvolvimento do front-end;
- Referência de padrões tanto para os designers, como para os desenvolvedores.

Uma documentação completa e detalhada é essencial para garantir uma padronização de escolhas de design, além de transparência e alinhamento entre os integrantes do time, os stakeholders, e entre futuros designers e desenvolvedores que farão manutenção da solução.

Tudo isso foi feito pensando em boas práticas de UX (User Experiencie), seguindo padrões e diretrizes epecíficas que tem como foco a experiência do usuário no centro, como:

 - **Facilidade de Uso**: Garantir que a interface seja intuitiva e fácil de navegar, reduzindo a curva de aprendizado e permitindo que os usuários comecem a trabalhar com o mínimo de treinamento.
 - **Acessibilidade**: Design que é acessível a todos os usuários, incluindo aqueles com limitações visuais ou outras necessidades especiais.
 - **Eficiência**: Otimização do fluxo de trabalho do usuário para maximizar a produtividade, permitindo que tarefas comuns sejam executadas de maneira rápida e eficiente.


 Além disso, foram utilizadas outras diretrizes de UX que serão detalhadas no decorrer deste documento.
## 1.2 Como deve ser utilizado
### Acesso ao design system
Para acessar o design system completo do projeto, contendo todos os artefatos em sua completude, é necessário acessar o nosso [Figma](https://www.figma.com/design/p8cF4hzhWoL1W7mPtXXbWC/DeployBuddy?node-id=0%3A1&t=hRt9LdgaM27RrRWF-1).
### Design de novas features
Para a criação de novas features ou telas, é indicada a leitura destes documento, bem como uma familiarização com o Design System por completo na aba "Design System" do nosso figma, para que haja um entendimento sobre quais padrões devem ser seguidos, e quais componentes devem ser reutilizados, a fim de seguir a mesma identidade visual das features já desenvolvidas. Dessa forma, é possivel criar um ecossistema coeso, em que o usuário entende o sistema todo como algo consistente, entregando uma boa experiência nas interações com a plataforma.
### Codificação da interface
Para a codificação da interface - o FrontEnd - é importante que seja feito um setup inicial em um arquivo global de estilos, contendo os padrões definidos pelo Guia de Estilos. Além disso, os componentes da biblioteca de componentes devem ser criados em um diretório "components", para que possam ser reutilizados no decorrer do desenvolvimento. Para isso, é possivel visualizar cada componente como código .css, e replicar com os ajustes necessários para a aplicação.
### Manutenção e atualização
- Consistência: Sempre verifique se novos componentes ou estilos seguem os padrões estabelecidos.
- Documentação: Atualize esta documentação sempre que houver mudanças ou adições significativas ao design system. ë importante interpretarmos este documento como algo dinâmico, que está sendo passando por atualizações e melhorias constantes! Isso é essencial para uma solução que está em constante desenvolvimento e melhoria.
- Sugestão de melhoria futura: Pensando em escalabilidade, seria interessante disponibilizar a biblioteca de componentes como uma biblioteca a ser importada em frameworks de frontend como react.js - o que estamos utilizando no momento. Isso facilitaria o desenvolvimento e aumentaria a escalabilidade do design system. Como este é um MVP, optamos por não fazer esse desenvolvimento nessa etapa incial, visando uma otimização de tempo e recursos.
## 1.3 Referências
Nosso design system foi inspirado por algumas das melhores práticas de design de sistemas disponíveis no mercado. Uma das principais referências que utilizamos foi o IBM Carbon Design System. O Carbon Design System é conhecido por sua robustez, flexibilidade e a atenção detalhada aos princípios de design, que ajudaram a guiar a estruturação e os padrões adotados no nosso sistema.

IBM. Carbon Design System. Disponível em: https://carbondesignsystem.com. Acesso em: 26 maio 2024.

## 2. Validação do Design System
Com base no desenvolvimento do Design System, estruturamos uma abordagem didática para avaliar os artefatos com o parceiro, com o objetivo de validação de ideias, alinhamento de expectativas e aprimoramento contínuo do sistema. Esta validação ocorreu durante a Sprint Review (da sprint 3), com o representante Bruno, da Everymind, onde foram coletados os seguintes feedbacks:

### 2.1. Feedbacks coletados
- Validação da Consistência Visual: A definição e aplicação correta das diretrizes visuais foram validadas. O representante destacou que a consistência no uso das cores, tipografia e espaçamento contribui significativamente para uma experiência de usuário coerente e agradável.

- Inclusão de Novos Componentes: Houve uma sugestão de inclusão de novos componentes específicos para interações mais complexas, como modais dinâmicos e tabelas interativas, que são cruciais para a operação diária dos usuários.

- Biblioteca de Padrões Eficiente: Foi elogiada a eficiência da biblioteca de padrões em garantir uma interface intuitiva e uma navegação fluida. A padronização de elementos comuns foi vista como uma maneira eficaz de reduzir a curva de aprendizado.

- Melhorias na Acessibilidade: Houve feedbacks sobre a necessidade de aumentar o contraste de alguns elementos para melhorar a acessibilidade, além de garantir que todos os componentes fossem totalmente navegáveis via teclado e compatíveis com leitores de tela.


### 2.2. Insights e melhorias
Com base nos feedbacks recebidos, refletimos sobre diversas oportunidades de melhoria nos artefatos desenvolvidos. Essa reflexão gerou insights valiosos e contribuiu para um melhor entendimento dos fluxos e requisitos da solução. Alguns dos insights são:

- Foco na Inclusão de Novos Componentes: Incluir componentes adicionais que suportem interações mais complexas, como modais dinâmicos e tabelas interativas, para atender melhor às necessidades dos usuários.
- Aprimorar a Acessibilidade: Revisar e melhorar os contrastes de cores e garantir que todos os componentes sejam navegáveis via teclado e compatíveis com leitores de tela.
- Personalização da Interface: Desenvolver opções de customização de temas que permitam maior flexibilidade e personalização da interface pelos usuários.
- Engajamento Contínuo com os Usuários: Continuar utilizando os materiais de entendimento do usuário de forma ativa no desenvolvimento diário, incorporando feedbacks e ajustando o Design System conforme necessário.

## 3. Guia de Estilo

Segue o guia de estilo visual da Deploy Buddy, que constitui um componente essencial para a manutenção da consistência e identidade visual em toda a plataforma. Este guia abrange especificações detalhadas sobre as cores, tipografia e iconografia utilizadas na interface do usuário. Ele serve como uma referência crucial para designers e desenvolvedores, assegurando que todos os elementos visuais estejam alinhados com a estética e os padrões estabelecidos para a solução Deploy Buddy.

![styleguide](../assets/styleguidemais.png)

### 3.1 Grid 

O conceito de grid em design de interfaces é essencial para criar layouts organizados e visualmente atraentes. Utilizando uma estrutura composta de linhas horizontais e/ou verticais que se intersectam, o grid permite a disposição coerente de elementos em uma interface, melhorando tanto a estética quanto a funcionalidade do design.

Um exemplo comum e eficaz de sistema de grid é o de 12 colunas, que foi o escolhido pelo grupo para a utilização nesse projeto, muito utilizado por sua flexibilidade e versatilidade. Este sistema divide a área de layout em 12 colunas iguais, permitindo que designers organizem o conteúdo em várias configurações que se adaptam a diferentes tamanhos de tela. Essa abordagem facilita o design responsivo, garantindo que o layout seja eficiente e acessível em dispositivos variados. 

A imagem abaixo ilustra como o grid de 12 colunas é aplicado na prática, neste caso,  na tela de kanban que funciona para fazer o gerenciamento de deploy. Este exemplo mostra como o grid foi utilizado para segmentar e organizar visualmente as etapas de deploy de software, facilitando a leitura e o gerenciamento das informações em diferentes fases do processo.

### 3.1 Diretrizes de Design

As diretrizes de design foram desenvolvidas para suportar a consistência visual e funcional em toda a plataforma. Estas incluem:

**Layout Consistente**: Os elementos de interface são organizados de forma consistente nas diversas páginas, promovendo uma experiência de usuário coesa.

**Paleta de Cores**: Utilização de uma paleta de cores suave, mas distintiva, que reflete a modernidade e é agradável aos olhos, minimizando a fadiga visual em uso prolongado.

**Tipografia**: Emprego de tipografia clara e legível, com tamanhos de fonte adequados para leitura e hierarquia visual bem definida.

## 4. Biblioteca de padrões

Nossa biblioteca de padrões é um repositório organizado de componentes reutilizáveis e práticas recomendadas de design e desenvolvimento. Ela nos ajuda a garantir a consistência, eficiência e qualidade do software. Nesta biblioteca, você encontra componentes de interface do usuário (UI) reutilizáveis, layouts e exemplos de hierarquia visual na interface. Alguns desses padrões vêm acompanhados de exemplos corretos e incorretos de uso, enquanto outros não possuem esses exemplos, pois esses conceitos são abordados de outra forma em outras partes do sistema de design.

Abaixo segue o detalhamento de cada seção da biblioteca de componentes e sua descrição. É recomendado o acompanhamento visual do Figma para maior entendimento dessa documentação.


![styleguide](../assets/DeployBuddy_bibli_padro.png)


### 4.1 Hierarquia visual - Textos

A hierarquia vertical em textos é fundamental para garantir que o usuário leia primeiro as informações mais importantes. Essa estrutura organiza o conteúdo de forma que o item de maior relevância seja destacado no topo, seguido por informações adicionais em ordem de importância decrescente.

Em um modal de alerta, o título deve ser claro e direto, geralmente informando sobre a ação que o usuario deve tomar, como "Alerta de conflito". Logo abaixo, deve haver uma breve descrição ou mensagem que forneça mais detalhes sobre o alerta e suas possíveis ações.

Já a hierarquia de tamanho e contraste é crucial para definir a importância das informações e orientar a atenção do usuário. Esse princípio é amplamente utilizado na seção introdutória das páginas, onde é essencial captar o interesse do usuário rapidamente e transmitir as informações-chave de forma eficaz.

Elementos de maior tamanho e contraste, como títulos, devem ser usados para destacar as informações mais importantes. Por exemplo, o título principal da página deve ser grande e visualmente distinto, atraindo imediatamente os olhos do usuário. Texto de apoio, embora menores, ainda devem ser suficientemente contrastantes para serem facilmente legíveis e para guiar o fluxo de leitura.

### 4.2 Hierarquia visual - Contraste de cores

O contraste de cores desempenha um papel essencial na interface do usuário, evidenciando elementos selecionados pelo usuário ou definidos pelo sistema. Essa técnica é vital para proporcionar uma visão clara das ações que estão ocorrendo e das opções disponíveis para o usuário, melhorando significativamente a experiência de navegação e interação.

Deve ser feito o uso de cores contrastantes para destacar elementos selecionados permite que o usuário identifique rapidamente quais itens estão ativos ou em foco. Esse contraste não só facilita a localização visual imediata, mas também reforça a percepção de interação, indicando ao usuário que sua ação foi registrada pelo sistema.

### 4.3 Hierarquia visual - White Space

Estabelecer hierarquia em componentes utilizando espaços em branco (white space) entre elementos é uma prática fundamental para criar uma interface de usuário clara e organizada. O uso estratégico do espaço em branco não apenas melhora a estética do design, mas também orienta a atenção do usuário e facilita a compreensão das informações apresentadas.

O espaço em branco ajuda a separar visualmente diferentes seções e componentes da interface, criando uma hierarquia natural. Elementos importantes, como títulos e botões, devem ser destacados ao serem cercados por mais espaço em branco, fazendo com que se destaquem e atraiam mais atenção.

### 4.4 Hierarquia visual - Cores

As cores dos botões e componentes são extremamente importantes, especialmente quando a cor do texto e dos ícones precisa estar na mesma paleta do botão. Esse princípio é frequentemente utilizado para elementos que precisam ser diferenciados em conjuntos e para botões que representam escolhas críticas, como ignorar ou resolver um conflito de merge.

A harmonia de cores entre o texto, ícones e botões assegura uma interface visualmente coerente e facilita a identificação e compreensão dos elementos. Por exemplo, um botão de ação principal, como "Resolver Conflito", deve ter uma cor distintiva que o destaque dos demais botões. O texto e os ícones nesse botão devem ser ajustados para garantir legibilidade e contraste adequados, utilizando tons da mesma paleta de cores.

Quando dois botões estão agrupados, a consistência de cores precisa ajudar a diferenciar e categorizar visualmente as opções. Botões que representam ações opostas, como "Resolver manualmente" e "Aplicar Solução", devem ter cores contrastantes para indicar claramente suas funções diferentes. Por exemplo, "Resolver manualmente" deve ser apresentado em uma cor negativa, de vermelho, enquanto "Aplicar Solução" deve ter uma cor positiva e chamativa, a verde.

### 4.5 Componentes de UI Reutilizáveis - Layout

Um layout reutilizável com navbar e sidebar, com margens e paddings definidos, é essencial para criar uma interface de usuário consistente e eficiente. Esse tipo de layout garante que o conteúdo seja apresentado de forma organizada, facilitando a navegação e a interação do usuário com o sistema.

**1. Navbar (Barra de Navegação)**: A navbar posicionada no topo da página, está presente em todas as páginas sendo dinâmica de acordo com o usúario que esta logado, por isso está presente no layout.

**2. Sidebar (Barra Lateral)**: A sidebar, que fica à esquerda das páginas, assim como a navbar, também é presente em todos as páginas e é necessário que esteja contida no layout padrão.

**4. Margins e Paddings no Layout Principal**: Para garantir uma experiência de usuário fluida e coerente, o layout principal deve considerar paddings bem definidos, espaçamento interno do conteúdo principal para melhorar a legibilidade e a organização dos elementos.

### 4.6 Componentes de UI Reutilizáveis - Botões

Botões são elementos reutilizáveis importantes para nossa aplicação, oferecendo consistência e eficiência na interação do usuário. Embora sejam essencialmente semelhantes, os botões variam em estilo e função para atender a diferentes necessidades, como botões com ou sem preenchimento (fill), botões de destaque com cores para ações importantes, e botões com ícones que facilitam a usabilidade.

**1. Botões com Preenchimento (Fill)**:
Botões com preenchimento são usados para destacar ações primárias. Eles são visualmente proeminentes e atraem a atenção do usuário, tornando-os ideais para ações principais que requerem uma resposta imediata, como "Aplicar Solução".

**2. Botões sem Preenchimento (Outline)**: Botões sem preenchimento têm apenas uma borda delineada e são visualmente menos proeminentes que os botões preenchidos. Eles são usados para ações secundárias ou menos críticas, ajudando a manter o foco do usuário nas ações primárias sem adicionar muita distração.

**3. Botões de Destaque para Ações Importantes**: Para ações críticas ou de alta prioridade, botões de destaque são essenciais. Esses botões devem ser preenchidos com cores vibrantes para indicar sua importância. Por exemplo, botões como "Resolver Manualmente" ou "Aplicar Soluções" precisam ser destacados para que o usuário os identifique rapidamente e compreenda sua importância.

**4. Botões com Ícones**: Botões com ícones incorporados são usados para facilitar a usabilidade e transmitir a funcionalidade de maneira visual. Ícones ajudam a comunicar a ação do botão de forma rápida e eficaz, complementando o texto ou, em alguns casos, substituindo-o completamente em interfaces onde o espaço é limitado.

### 4.7 Componentes de UI Reutilizáveis - Modais Gerais

Usamos modais fundamentalmente para criar projetos e organizações em nossa plataforma e para alertar os usuários. Eles seguem uma estrutura simples e eficaz, composta por: um título com um ícone de fechar modal no canto, uma linha horizontal separando o título do conteúdo subsequente, o conteúdo principal e os botões de ação.

**1. Título e Ícone de Fechar**: No topo do modal, o título deve ser claro e conciso, refletindo a finalidade do modal, como "Criar Novo Projeto" ou "Alerta de conflito". O ícone de fechar modal, localizado no canto superior direito, permite que os usuários fechem o modal facilmente se decidirem não prosseguir com a ação. Este ícone deve usar um "X" para garantir acessibilidade e usabilidade e ser um design familiar.

**2. Linha Horizontal de Separação**: Uma linha horizontal logo abaixo do título serve para separar visualmente o cabeçalho do conteúdo principal. Esta linha ajuda a organizar o layout, criando uma distinção clara entre o título e o restante do modal, melhorando a legibilidade e a estrutura visual.

**3. Conteúdo Principal**: O conteúdo do modal varia de acordo com a sua finalidade. Para modais de criação de projetos e organizações, o conteúdo inclui formulários e campos de texto. Em modais de alerta, o conteúdo geralmente consiste em mensagens curtas e diretas que informam o usuário sobre uma ação ou situação específica, como erros, avisos ou confirmações.

**4. Botões de Ação**: Na parte inferior do modal, botões de ação permitem ao usuário confirmar ou cancelar a operação. Botões como "Cancelar Merge" e "Visualizar Sugestão" devem ser claramente identificados e posicionados de forma lógica. O botão de ação primária, que executa a principal função do modal, deve ser visualmente destacado (fill), enquanto o botão de ação secundária deve ser menos proeminente (outline).

### 4.8 Componentes de UI Reutilizáveis - Modais de Feedbacks

Os modais de feedback são utilizados para informar aos usuários o resultado de suas ações na plataforma. Eles são divididos em três tipos: vermelho para falha, verde para sucesso e amarelo para atenção. Enquanto os modais vermelhos e verdes fornecem feedback após a realização de uma ação, o modal amarelo serve para alertar durante o processamento de uma ação ou para confirmar o impacto potencial de uma ação específica.

**1. Modal Vermelho (Falha)**: Os modais vermelhos são usados para indicar falhas ou erros. Eles são visualmente impactantes e chamativos, utilizando o vermelho como cor principal para chamar a atenção imediata do usuário.

**2. Modal Verde (Sucesso)**: Os modais verdes são usados para indicar o sucesso de uma ação. Eles transmitem uma sensação de realização e confirmação positiva.

**3. Modal Amarelo (Atenção)**: Os modais amarelos são usados para alertar o usuário durante o processamento de uma ação ou para confirmar ações que podem ter consequências importantes.

### 4.9 Design Responsivo

Estamos aplicando regras de design responsivo para garantir que o conteúdo apareça de maneira apropriada em cada dispositivo, independentemente da resolução. Essas regras determinam como o conteúdo flui e como o layout muda, adaptando-se automaticamente ao tamanho disponível. Utilizamos três regras principais em nosso projeto para assegurar que o conteúdo se reorganize dinamicamente conforme necessário: grids fluidas, imagens fluidas e media queries.

**1. Grid Fluida**: As grids fluidas permitem que o layout se ajuste automaticamente ao tamanho da tela do dispositivo. À medida que a tela redimensiona, as colunas da grid fluida se expandem ou contraem proporcionalmente, garantindo que o conteúdo permaneça organizado e legível, fazemos isso através da utilização de unidades de medida relativas, como porcentagens, em vez de unidades fixas, como pixels, para definir a largura das colunas e áreas do layout.

**2. Imagens Fluidas**: As imagens fluidas garantem que os elementos visuais se ajustem adequadamente ao tamanho da tela, mantendo a proporção e a qualidade. Fazemos isso definindo as imagens com largura máxima de 100% e altura automática para que se ajustem dinamicamente ao container em que estão inseridas.

**2. Media Queries**: As media queries permitem a aplicação de estilos específicos para diferentes tamanhos de tela e resoluções, garantindo que o layout se adapte a diversos dispositivos. Fazemos isso verificando o tamanho da tela do dispositivo e aplicando os estilos correspondentes, ajustando o layout, tipografia, espaçamento e outros elementos de design.

![Media Querias](../assets/media_queries.png)

## 5. Biblioteca de Componentes

No projeto **Deploy Buddy**, utilizamos componentes em React para desenvolver a interface de usuário. Este framework nos permite criar componentes reutilizáveis e modulares, garantindo que nossos padrões de design e usabilidade sejam mantidos de maneira consistente em toda a aplicação.

![biblioteca de componentes](../assets/complib.png)

### Especificidades dos Componentes

Cada componente da biblioteca foi projetado com base em princípios de design moderno e práticas recomendadas para garantir usabilidade, acessibilidade e eficiência.

1. **Consistência Visual**: Todos os componentes seguem uma paleta de cores e tipografia unificadas para proporcionar uma experiência de usuário coerente e profissional.
2. **Reutilização**: Componentes como botões, cards e formulários são projetados para serem reutilizados em diferentes partes da aplicação, reduzindo o esforço de desenvolvimento e manutenção.
3. **Responsividade**: Os componentes são adaptáveis a diferentes tamanhos de tela, garantindo uma experiência ótima em dispositivos móveis e desktops.
4. **Acessibilidade**: Todos os componentes são desenvolvidos considerando boas práticas de acessibilidade, como contraste de cores adequado e suporte para navegação por teclado.

## 6. Prototipação

![Tela de menu](../assets/prototypemenu.png)
![Tela de kanban](../assets/prototypekanban-page.png)
![Tela de alert](../assets/prototypealert.png)

Para uma visão mais detalhada acessar o arquivo Figma completo [clicando aqui.](https://www.figma.com/design/p8cF4hzhWoL1W7mPtXXbWC/DeployBuddy?node-id=0%3A1&t=hRt9LdgaM27RrRWF-1)

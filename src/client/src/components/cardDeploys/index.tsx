import React from "react";
import {
  Container,
  ContainerCards,
  Separator,
  Title,
  LastProjects,
  Message,
} from "./styles";
import SubCardDeploy from "../subCardDeploy";

function CardDeploys() {
  return (
    <Container>
      <Separator>
        <Title>
          <LastProjects>Últimos deploys</LastProjects>
          <Message>Projetos em que sou colaborador</Message>
        </Title>
      </Separator>
      <ContainerCards>
        <SubCardDeploy
          title="Criação de nova lista de usuários"
          date="20/05/2024 - 12:00"
          author="Marselo"
          env="prod"
        />
        <SubCardDeploy
          title="Outro Projeto"
          date="15/05/2024 - 10:00"
          author="Matheus"
          env="dev"
        />
        <SubCardDeploy
          title="Mais um Projeto"
          date="16/05/2024 - 10:00"
          author="EuroPapo Hagge"
          env="uat"
        />
      </ContainerCards>
    </Container>
  );
}

export default CardDeploys;

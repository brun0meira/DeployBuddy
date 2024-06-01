import React, { useState } from "react";
import { AddButton, Container, ContainerCards, Separator, Title, LastProjects, Message} from "./style";
import CardDashboard from "..//cardDashboard";
import ModalAddProject from "../modalAddProject";

function BackgroundCards() {

  const [isModalVisible, setModalVisible] = useState(false);

  return (
    <Container>
    <Separator>
      <Title>
        <LastProjects>Últimos Projetos</LastProjects>
        <Message>Projetos em que sou colaborador</Message>
      </Title>
    <div>
    <AddButton onClick={() => setModalVisible(true)}>+ Adicionar Projeto</AddButton>
    <ModalAddProject isVisible={isModalVisible} onClose={() => setModalVisible(false)} />
    </div>
    </Separator>
    <ContainerCards>
        <CardDashboard></CardDashboard> 
        <CardDashboard></CardDashboard>
    </ContainerCards>
    </Container>
  );
}

export default BackgroundCards;

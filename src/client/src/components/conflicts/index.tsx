// ConflictsModal.tsx
import React from "react";
import {
  Backdrop,
  ModalContainer,
  Header,
  Title,
  CloseButton,
  ContentAlert,
  ButtonContainer
} from "./styles";
import Button from "../buttonOutlined/buttonOutlined";

type Props = {
  isVisible: boolean;
  onClose: () => void;
};

const ConflictsModal: React.FC<Props> = ({ isVisible, onClose }) => {
  if (!isVisible) return null;

  return (
    <Backdrop onClick={onClose}>
      <ModalContainer onClick={(e) => e.stopPropagation()}>
        <Header>
          <Title>Alerta de Conflito!</Title>
          <CloseButton onClick={onClose}>X</CloseButton>
        </Header>
        <ContentAlert>
          Houve um conflito na implantação de ”release #13” no ambiente
          “Desenvolvimento”. Mas não se preocupe, o Merge Solver tem uma
          sugestão de solução pra você!
        </ContentAlert>
        <ButtonContainer>
          <Button
            onClick={() => alert("Outline Button Clicked!")}
            variant="outline"
            size="small"
          >
            Cancelar Merge
          </Button>
          <Button
            onClick={() => alert("Outline Button Clicked!")}
            variant="secondary"
            size="small"
          >
            Visualizar Sugestão
          </Button>
        </ButtonContainer>
      </ModalContainer>
    </Backdrop>
  );
};

export default ConflictsModal;

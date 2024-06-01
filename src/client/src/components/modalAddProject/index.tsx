// index.tsx
import React, { useState } from 'react';
import { Backdrop, ModalContainer, Fields, Inputs, SaveButton, CloseButton, Title, Header } from './style';

type Props = {
  isVisible: boolean;
  onClose: () => void;
};

const ModalAddProject: React.FC<Props> = ({ isVisible, onClose }) => {
  if (!isVisible) return null;

  return (
    <Backdrop onClick={onClose}>
      <ModalContainer onClick={e => e.stopPropagation()}>
        <Header>
          <Title>Adicionar projeto</Title>
          <CloseButton onClick={onClose}>X</CloseButton>
        </Header>
        <form>
          <Fields>
            Nome:
            <Inputs type="text" placeholder="Nome do projeto" />
            String de conexão:
            <Inputs type="text" placeholder="String de conexão Salesforce" />
          </Fields>
          <div style={{ display: 'flex', justifyContent: 'flex-end', marginTop: '20px' }}>
            <SaveButton type="submit">Adicionar Projeto</SaveButton>
          </div>
        </form>
      </ModalContainer>
    </Backdrop>
  );
};

export default ModalAddProject;

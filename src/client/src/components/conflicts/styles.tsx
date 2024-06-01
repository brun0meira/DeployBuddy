import styled from "styled-components";

export const Backdrop = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
`;

export const ModalContainer = styled.div`
  background: white;
  padding: 20px;
  border-radius: 5px;
  width: 40%;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
`;

export const Header = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  border-bottom: 1px solid #ccc;
`;

export const Title = styled.p`
  flex-grow: 1;
  text-align: center;
  margin: 0;
  font-size: 20px;
  margin-bottom: 20px;
`;

export const CloseButton = styled.button`
  border: none;
  background: none;
  cursor: pointer;
  font-size: 16px;
  margin-bottom: 15px;
  margin-left: 5px;
`;

export const ContentAlert = styled.h3`
  text-align: center;
  margin: 10px;
`;

export const ButtonContainer = styled.div`
  display: flex;
  justify-content: center;
  gap: 10px;
  margin-top: 20px;
`;

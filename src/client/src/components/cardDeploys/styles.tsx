// carddeploys style

import styled from "styled-components";

export const Container = styled.div`
  border: 1px solid #ccc;
  padding: 20px;
  margin: 20px 0;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  background-color: #fff;
  display: flex;
  flex-direction: column;
  width: 50%;
`;

export const ContainerCards = styled.div`
  display: flex;
  flex-direction: column;  /* Alterado para coluna */
`;

export const Title = styled.div`
`;

export const LastProjects = styled.h1`
  font-size: 18px;
  color: #333;
  margin: 0 0 10px 0;
`;

export const Message = styled.p`
  margin: 0;
  color: #666;
  font-size: 12px;
`;

export const Separator = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`;

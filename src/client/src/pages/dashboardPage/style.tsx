//dashboard page style
import styled from "styled-components";

export const MainContainer = styled.div`
  display: flex;
`;

export const ContentContainer = styled.div`
  flex-grow: 1; 
  display: flex;
  flex-direction: column; 
`;

export const PageContent = styled.div`
  flex-grow: 1; 
  padding: 20px;
`;

export const CardsBackground = styled.div`
  border: 1px solid #ccc; 
  padding: 20px; 
  margin: 20px 0;
  border-radius: 5px; 
  box-shadow: 0 2px 5px rgba(0,0,0,0.1); 
  background-color: #fff; 
  display: flex;
  flex-direction: row;
`;
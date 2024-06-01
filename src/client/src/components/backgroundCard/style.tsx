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
`;

export const ContainerCards = styled.div`
  display: flex;
  flex-direction: row;
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

export const AddButton = styled.button`
  padding: 8px 8px; 
  border-radius: 4px; 
  border: none; 
  width: 160px;
  background-color: rgba(214, 224, 234, 0.35); 
  color: white; 
  cursor: pointer;
  font-size: 14px;
  color:  #00458E; 
  align-self: center;
`;

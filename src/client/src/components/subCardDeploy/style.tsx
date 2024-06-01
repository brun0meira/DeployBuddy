// subcard deploy style

import styled from "styled-components";

export const Card = styled.div`
  display: flex;
  flex-direction: column;
  border-radius: 5px;
  padding: 10px;
  margin-bottom: 10px;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.1);
  background-color: #fff;
`;

export const CardHeader = styled.div`
  display: flex;
  align-items: center;
`;

export const CardIcon = styled.div`
  margin-right: 10px;
`;

export const CardContent = styled.div`
  display: flex;
  flex-direction: column;
`;

export const CardText = styled.h2`
  font-size: 16px;
  color: #333;
  margin: 0;
`;

export const CardDetails = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center; /* Certifica que todos os elementos estão alinhados */
  margin-top: 5px;
`;

export const CardDate = styled.p`
  margin: 0;
  color: #666;
  font-size: 12px;
`;

export const CardInfo = styled.p`
  margin: 0;
  color: #666;
  font-size: 12px;
  margin-left: 15px;
`;

export const CardEnv = styled.div`
  display: flex;
  align-items: center;
  margin-left: auto;
  justify-content: space-between;
  
`;

export const CardEnvIcon = styled.div`
  margin-right: 5px;
`;

export const CardEnvText = styled.p`
  margin: 0;
  color: #666;
  font-size: 12px;
`;

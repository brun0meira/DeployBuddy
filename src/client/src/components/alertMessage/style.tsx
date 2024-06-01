import styled from 'styled-components';

interface AlertContainerProps {
    type: 'success' | 'error';
  }

  export const AlertContainer = styled.div.attrs<AlertContainerProps>(({ type }) => ({
    type,
  }))<AlertContainerProps>`
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #fff;
    padding: 15px 20px;
    border-radius: 12px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    margin: 10px 0;
  `;

  export const AlertMessage = styled.p.attrs<AlertContainerProps>(({ type }) => ({
    type,
  }))<AlertContainerProps>`
    color: ${({ type }) => (type === 'success' ? '#4CAF50' : '#DA1E28')};
    font-size: 16px;
    margin: 0;
  `;

export const AlertLink = styled.a`
  color: #aaa;
  text-decoration: none;
  font-size: 14px;
  &:hover {
    text-decoration: underline;
  }
`;
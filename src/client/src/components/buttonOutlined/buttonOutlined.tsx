// Button.tsx
import React from 'react';
import styled from 'styled-components';

interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  variant?: 'primary' | 'secondary' | 'outline';
  size?: 'small' | 'medium' | 'large';
}

const ButtonStyled = styled.button<ButtonProps>`
  display: inline-block;
  padding: ${(props) =>
    props.size === 'small' ? '8px 16px' :
    props.size === 'medium' ? '12px 24px' :
    '16px 32px'};
  font-size: ${(props) =>
    props.size === 'small' ? '12px' :
    props.size === 'medium' ? '14px' :
    '16px'};
  color: ${(props) =>
    props.variant === 'primary' ? '#fff' :
    props.variant === 'secondary' || props.variant === 'outline' ? '#00458E' :
    '#007bff'};
    
  background-color: ${(props) =>
    props.variant === 'primary' ? '#E8EDF3' :
    props.variant === 'secondary' ? '#D6E0EA' :
    'transparent'};
  border: ${(props) =>
    props.variant === 'outline' ? '2px solid #00458E' :
    'none'};
  border-radius: 4px;
  cursor: pointer;
  text-align: center;
  transition: background-color 0.3s ease, color 0.3s ease;
  &:hover {
    background-color: ${(props) =>
      props.variant === 'primary' ? '#0056b3' :
      props.variant === 'secondary' ? '#0056b3' :
      '#e2e6ea'};
    color: ${(props) =>
      props.variant === 'outline' ? '#0056b3' : '#fff'};
  }
`;

const Button: React.FC<ButtonProps> = ({ children, onClick, variant = 'primary', size = 'medium' }) => {
  return (
    <ButtonStyled onClick={onClick} variant={variant} size={size}>
      {children}
    </ButtonStyled>
  );
};

export default Button;

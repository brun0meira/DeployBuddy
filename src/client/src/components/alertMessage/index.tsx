import React, { useEffect } from 'react';
import { AlertContainer, AlertMessage, AlertLink } from './style';

interface AlertsMessageProps {
  type: 'success' | 'error';
  message: string;
  onViewDetails?: () => void;
  onClose?: () => void; // Adiciona uma função opcional para fechar o alerta
  duration?: number; // Adiciona uma duração opcional em milissegundos
}

const AlertsMessage: React.FC<AlertsMessageProps> = ({ type, message, onViewDetails, onClose, duration = 5000 }) => {
  useEffect(() => {
    if (onClose) {
      const timer = setTimeout(() => {
        onClose();
      }, duration);

    
      return () => clearTimeout(timer);
    }
  }, [onClose, duration]);

  return (
    <AlertContainer type={type}>
      <AlertMessage type={type}>{message}</AlertMessage>
      {type === 'error' && (
        <AlertLink href="#" onClick={onViewDetails}>Visualizar detalhes</AlertLink>
      )}
    </AlertContainer>
  );
};

export default AlertsMessage;
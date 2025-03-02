import React, { useState } from "react";
import {
  Container,
  LoginForm,
  Input,
  LoginButton,
  LogoStyle,
  LoginText,
  TitleLogin,
  ErrorMessage,
} from "./LoginStyles";
import LogoDeployBuddy from "../../assets/Logo";
import { useNavigate } from 'react-router-dom';  

function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    if (email.endsWith("@everymind.com")) {
      console.log("Submit:", email, password);
      setError("");
      let user = fetch('http://localhost:8080/api/v1/users/auth', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })
        .then((response) => response.json())
        .then((data) => {
          console.log('Success:', data);
          localStorage.setItem('token', data);
          navigate('/Dashboard');
        })
        .catch((error) => {
          console.error('Error:', error);
      });

      // navigate('/Dashboard');
    } else {
      setError("Please enter a valid Everymind email.");
    }
  };

  return (
    <Container>
      <LoginForm>
        <LogoStyle>
          <LogoDeployBuddy/>
        </LogoStyle>
        <TitleLogin>
          Sign in to DeployBuddy with your Everymind email.
        </TitleLogin>
        <form onSubmit={handleSubmit}>
          <LoginText>Email address:</LoginText>
          {error && <ErrorMessage>{error}</ErrorMessage>}
          <Input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <LoginText>Password:</LoginText>
          <Input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <LoginButton type="submit">Login</LoginButton>
        </form>
      </LoginForm>
    </Container>
  );
}

export default LoginPage;

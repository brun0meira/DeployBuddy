import React from "react";
import {
  Card,
  CardContent,
  CardHeader,
  CardInfo,
  CardDate,
  CardIcon,
  CardText,
  CardDetails,
  CardEnv,
  CardEnvText,
  CardEnvIcon,
} from "./style";
import {
  FileUp,
  GitPullRequestArrow,
  FileTerminal,
  MonitorDot,
} from "lucide-react";

interface SubCardDeployProps {
  title: string;
  date: string;
  author: string;
  env: "prod" | "uat" | "dev";
}

const SubCardDeploy: React.FC<SubCardDeployProps> = ({
  title,
  date,
  author,
  env,
}) => {
  const renderEnvIcon = () => {
    switch (env) {
      case "prod":
        return <GitPullRequestArrow color="#C9CCED" size={24} />;
      case "uat":
        return <MonitorDot color="#BAE7EC" size={24} />;
      case "dev":
        return <FileTerminal color="#FFE6B5" size={24} />;
      default:
        return null;
    }
  };

  const renderEnvText = () => {
    switch (env) {
      case "prod":
        return "Produção";
      case "uat":
        return "UAT";
      case "dev":
        return "Development";
      default:
        return "";
    }
  };

  return (
    <Card>
      <CardHeader>
        <CardIcon>
          <FileUp color="#00458E" size={32} />
        </CardIcon>
        <CardContent>
          <CardText>{title}</CardText>
          <CardDetails>
            <CardDate>{date}</CardDate>
            <CardInfo>Autor: {author}</CardInfo>
          </CardDetails>
        </CardContent>
        <CardEnv>
          <CardEnvIcon>{renderEnvIcon()}</CardEnvIcon>
          <CardEnvText>{renderEnvText()}</CardEnvText>
        </CardEnv>
      </CardHeader>
    </Card>
  );
};

export default SubCardDeploy;

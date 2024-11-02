import {
  Link,
  isRouteErrorResponse,
  useLocation,
  useRouteError,
} from "@remix-run/react";
import { Button } from "../ui/button";

interface Props {
  retryPath?: string;
  children: React.ReactNode;
}

const ErrorCard = ({ retryPath, children }: Props) => (
  <div className="my-4 flex w-full flex-col">
    <div className="flex w-4/5 flex-col items-center justify-center self-center rounded-md bg-destructive p-2">
      <h1 className="line-clamp-1 text-ellipsis text-sm text-destructive-foreground">
        {children}
      </h1>
    </div>

    {retryPath && (
      <div className="mt-4 flex w-full flex-row justify-around">
        <Link className="w-1/3 self-center" to={retryPath}>
          <Button className="w-full min-w-min">Reintentar</Button>
        </Link>
        <Link className="w-1/3 self-center" to={"/"}>
          <Button className="w-full min-w-min">Volver a Home</Button>
        </Link>
      </div>
    )}
  </div>
);

export function Errors() {
  const error = useRouteError();
  const location = useLocation();

  if (isRouteErrorResponse(error)) {
    return (
      <ErrorCard retryPath={location.pathname}>
        {error.status} {error.statusText}
      </ErrorCard>
    );
  } else if (error instanceof Error) {
    return <ErrorCard retryPath={location.pathname}>{error.message}</ErrorCard>;
  } else {
    return <ErrorCard retryPath={location.pathname}>Unknown Error</ErrorCard>;
  }
}

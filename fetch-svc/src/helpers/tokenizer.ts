import jwt, { JsonWebTokenError } from "jsonwebtoken";
import { TokenPayload } from "../interface";

export function generateToken(payload: TokenPayload) {
  const token = jwt.sign(payload, process.env.JWT_SECRET!);
  return token;
}

export function decodeToken(token: string): TokenPayload {
  token = token.replace("Bearer ", "");
  const payload = jwt.verify(
    token,
    process.env.JWT_SECRET as string
  ) as TokenPayload;
  return payload;
}

export { JsonWebTokenError as TokenError };

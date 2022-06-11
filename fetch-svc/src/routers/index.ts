import { Router } from "express";
import comodityRouter from "./commodity";

const router = Router();

router.use(comodityRouter);
export default router;

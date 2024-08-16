export type User = {
  id: string;
  name: string;
  surname: string;
  email: string;
  password: string;
  role: string;
  createdAt: Date;
  updatedAt: Date;
};

export enum UserRoles {
  Coordenador = "Coordenador",
  Gestor = "Gestor",
  Secretaria = "Secretaria",
}

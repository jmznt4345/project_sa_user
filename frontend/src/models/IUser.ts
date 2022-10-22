import { RoleInterface } from "./IRole";

export interface UsersInterface {
    ID?: number,
    Name?: string;
    Email?: string;
    Phonenumber?: string;
    Password?: string;

    PositionID?: number;
    RoleID?: number;
    Role?: RoleInterface;
    GenderID?: number;
}
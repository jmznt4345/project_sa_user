import { RoleInterface } from "./IRole";

export interface UsersInterface {
    ID?: number,
    Name?: string;
    Email?: string;
    Phonenumber?: string;
    Password?: string;

    Educational_backgroundID?: number;
    RoleID?: number;
    Role?: RoleInterface;
    GenderID?: number;
}
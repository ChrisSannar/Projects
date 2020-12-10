import { getConnection } from 'typeorm';
import { User } from './entity/User';

export const resolvers = {
  Query: {
    getUser: async (_: any, args: any) => {
      const { id } = args

      return await getConnection().getRepository(User).findOne({ where: { id } });
    }
  },
  Mutation: {
    addUser: async (_: any, args: any) => {
      const { username, password } = args;
      try {
        const user = new User();
        user.username = username;
        user.password = password;

        await getConnection().getRepository(User).save(user);

        return true;
      } catch (error) {
        return false;
      }
    }
  }
}
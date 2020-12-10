
import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  ManyToMany,
  JoinTable
} from 'typeorm';

@Entity()
class Test {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  // The string value of the test
  @Column({ nullable: false })
  value: string;

  @ManyToMany(_ => Test2, (test: any) => test.value, {
    eager: true,    // Will always load the 'Test2' value with this one
    cascade: ['insert', 'update'],  // When you insert and update, but not delete
  })
  @JoinTable()
  others: Test2[]
}

@Entity()
class Test2 {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  // The string value of the test
  @Column({ nullable: false })
  value: string;

  // // Anywhere in the string we need to replace values (for repeated uses)
  // @ManyToMany(_ => Template, template => template.tag)
  // @JoinTable()
  // templateValues: Template[]
}

export default { Test, Test2 };
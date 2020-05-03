const index = require('./index.js');
const { List } = index;

describe('index.js', () => {
  it('should pull the index module', () => {
    expect(index).toBeDefined();
  });

  describe('Node List', () => {
    it('should be able to get list from the node modules', () => {
      const list = new List();
      expect(list).toBeDefined();
    });

    it('should add a new item to the list', () => {
      const listLength = 3;
      const list = setUpList(listLength);
      expect(list.length).toBe(listLength);
    });

    it('should clear', () => {
      const listLength = 3;
      const list = setUpList(listLength);
      expect(list.length).toBe(listLength);
      list.clear();
      expect(list.length).toBe(0);
      expect(list.head).toBeNull();
      expect(list.tail).toBeNull();
    });

    it('should find the item', () => {
      const list = setUpList(3);
      expect(list.find(item => item === 3)).toBe(3);
    });

  });
});

function setUpList(length) {
  let list = new List();
  let blankArr = Array.apply(null, Array(length)).map((item, index) => index + 1)
  for (let i of blankArr) {
    list.add(i);
  }
  return list;
}
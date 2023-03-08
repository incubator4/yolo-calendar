interface ICharacter {
  id: number;
  name: string;
  uid: number;
  live_id: number;
  main_color: string;
}

interface ICalendar {
  title: string;
  start_time: Date;
  end_time: Date;
  cid: number;
  tag_id: number;
}

interface ITag {
  id: number;
  name: string;
}

type CharacterCalendar = ICalendar & ICharacter;

interface IVtuber {
  id: number;
  name: string;
  uid: number;
  live_id: number;
  main_color: string;
}

interface ICalendar {
  id?: number;
  title: string;
  start_time: Date;
  end_time: Date;
  cid: number;
  tag_id: number;
  is_active: boolean;
}

interface ITag {
  id: number;
  name: string;
}

type VtuberCalendar = ICalendar & IVtuber;

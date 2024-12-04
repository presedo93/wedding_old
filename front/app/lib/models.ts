export interface Guest {
  id?: number;
  profile_id: string;
  created_at: string;
  updated_at: string;
  name: string;
  phone: string;
  is_vegetarian: boolean;
  allergies: string[];
  needs_transport: boolean;
}

export interface Profile {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
  phone: string;
  email: string;
  picture_url?: string;
  completed_profile: boolean;
  added_guests: boolean;
  added_songs: boolean;
  added_pictures: boolean;
}

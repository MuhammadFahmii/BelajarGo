# membuat folder utama
curr_date=$(date)
parent_folder="$1 at $curr_date"
mkdir "$parent_folder"

# membuat folder about me
about_me="$parent_folder/about_me"
mkdir "$about_me"

# membuat folder personal berisi facebook.txt
personal="$about_me/personal"
mkdir "$personal"
touch "$personal/facebook.txt"
echo "https://www.facebook.com/$2">"$personal/facebook.txt"

# membuat folder professional berisi linkedin.txt
professional="$about_me/professional"
mkdir "$professional"
touch "$professional/linkedin.txt"
echo "https://www.linkedin.com/in/$3">"$professional/linkedin.txt"

# membuat folder my_friend berisi list_of_my_friends.txt
my_friend="$parent_folder/my_friends"
mkdir "$my_friend"
touch "$my_friend/list_of_my_friends.txt"
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt>"$my_friend/list_of_my_friends.txt"

# membuat folder my_system_info berisi about_this_laptop.txt dan internet_connection.txt
my_system_info="$parent_folder/my_system_info"
mkdir "$my_system_info"
touch "$my_system_info/about_this_laptop.txt"
touch "$my_system_info/internet_connection.txt"

# isi txt 
echo "My Username $1">"$my_system_info/about_this_laptop.txt"
echo "With host $(whoami)">"$my_system_info/about_this_laptop.txt"
sudo ping -c 3 google.com>"$my_system_info/internet_connection.txt"
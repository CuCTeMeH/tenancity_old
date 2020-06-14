package user

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tenancity/api/core"
	gormUser "tenancity/api/user/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userID")
	var user gormUser.User
	// Raw SQL
	//core.Server.DB.Connections["fn"].Raw("select listings.id as username from `proxy_auctions` inner join `listings` on `proxy_auctions`.`listing_id` = `listings`.`id` and `proxy_auctions`.`deleted_at` is null left join `foreclosures_x_estates` on `foreclosures_x_estates`.`foreclosure_id` = `listings`.`foreclosure_id` and `foreclosures_x_estates`.`deleted_at` is null inner join `estates` on `foreclosures_x_estates`.`estate_id` = `estates`.`id` and `estates`.`deleted_at` is null left join `liens` on `estates`.`id` = `liens`.`estate_id` and `liens`.`deleted_at` is null left join `foreclosures` on `foreclosures`.`id` = `listings`.`foreclosure_id` and `foreclosures`.`deleted_at` is null where `proxy_auctions`.`deleted_at` is null and `listings`.`foreclosure_id` is not null and `proxy_auctions`.`status_id` in ('1', '2', '3', '6', '7', '5') and (`estates`.`case_city` in ('Addison', 'Algonquin', 'Alsip', 'Antioch', 'Aptakisic', 'Arbury Hills', 'Arlington Heights', 'Aurora', 'Bannockburn', 'Barrington', 'Barrington Hills', 'Barrington Woods', 'Bartlett', 'Batavia', 'Beach Park', 'Bedford Park', 'Beecher', 'Bellwood', 'Bensenville', 'Berkeley', 'Berwyn', 'Big Rock', 'Bloomingdale', 'Blue Island', 'Bolingbrook', 'Boulder Hill', 'Braidwood', 'Bridgeview', 'Broadview', 'Brookfield', 'Buffalo Grove', 'Bull Valley', 'Burbank', 'Burlington', 'Burnham', 'Burr Ridge', 'Burtons Bridge', 'Calumet City', 'Calumet Park', 'Campton Hills', 'Carol Stream', 'Carpentersville', 'Cary', 'Channahon', 'Chemung', 'Cherry Hill', 'Chicago', 'Chicago Heights', 'Chicago Ridge', 'Cicero', 'Clarendon Hills', 'Cloverdale', 'Country Club Hills', 'Countryside', 'Crest Hill', 'Crestwood', 'Crete', 'Crystal Lake', 'Darien', 'Deer Park', 'Deerfield', 'Del Mar Woods', 'Des Plaines', 'Diamond Lake', 'Dixmoor', 'Dolton', 'Downers Grove', 'East Dundee', 'East Hazel Crest', 'Eastwood Manor', 'Elburn', 'Elgin', 'Elk Grove Village', 'Elmhurst', 'Elmwood Park', 'Elwood', 'Evanston', 'Evergreen Park', 'Fairmont', 'Flossmoor', 'Ford Heights', 'Forest Lake', 'Forest Park', 'Forest View', 'Fox Lake', 'Fox Lake Hills', 'Fox River Grove', 'Frankfort', 'Franklin Park', 'Gages Lake', 'Geneva', 'Gilberts', 'Glen Ellyn', 'Glencoe', 'Glendale Heights', 'Glenview', 'Glenwood', 'Godley', 'Golf', 'Goodenow', 'Goodings Grove', 'Grandwood Park', 'Grayslake', 'Green Oaks', 'Greenwood', 'Gurnee', 'Hainesville', 'Half Day', 'Hampshire', 'Hanover Park', 'Hartland', 'Harvard', 'Harvey', 'Harwood Heights', 'Hastings', 'Hawthorn Woods', 'Hazel Crest', 'Hazel Green', 'Hebron', 'Hickory Hills', 'Highland Hills', 'Highland Park', 'Highwood', 'Hillside', 'Hinsdale', 'Hodgkins', 'Hoffman Estates', 'Holbrook', 'Holiday Hills', 'Homer Glen', 'Hometown', 'Homewood', 'Horatio Gardens', 'Huntley', 'Idylside', 'Indian Creek', 'Indian Head Park', 'Ingalls Park', 'Inverness', 'Island Lake', 'Itasca', 'Ivanhoe', 'Johnsburg', 'Joliet', 'Justice', 'Kaneville', 'Keeneyville', 'Kenilworth', 'Kildeer', 'Klondike', 'Knollwood', 'La Grange', 'La Grange Park', 'Lake Barrington', 'Lake Bluff', 'Lake Forest', 'Lake Villa', 'Lake Zurich', 'Lake in the Hills', 'Lakemoor', 'Lakewood', 'Lakewood Shores', 'Lansing', 'Lawrence', 'Lemont', 'Libertyville', 'Lily Cache', 'Lily Lake', 'Lilymoor', 'Lincoln Estates', 'Lincolnshire', 'Lincolnwood', 'Lindenhurst', 'Lisbon', 'Lisle', 'Lockport', 'Lombard', 'Long Grove', 'Long Lake', 'Lotus Woods', 'Lynwood', 'Lyons', 'Manhattan', 'Maple Park', 'Marengo', 'Markham', 'Marley', 'Matteson', 'Maywood', 'McCook', 'McCullom Lake', 'McHenry', 'Medinah', 'Melrose Park', 'Merrionette Park', 'Mettawa', 'Midlothian', 'Millbrook', 'Millburn', 'Millington', 'Mokena', 'Monee', 'Montgomery', 'Morton Grove', 'Mount Prospect', 'Mundelein', 'Naperville', 'New Lenox', 'Newark', 'Niles', 'Norridge', 'North Aurora', 'North Barrington', 'North Chicago', 'North Glen Ellyn', 'North Riverside', 'Northbrook', 'Northfield', 'Northlake', 'Oak Brook', 'Oak Forest', 'Oak Lawn', 'Oak Park', 'Oakbrook Terrace', 'Oakwood Hills', 'Old Mill Creek', 'Olympia Fields', 'Orland Hills', 'Orland Park', 'Oswego', 'Palatine', 'Palos Heights', 'Palos Hills', 'Palos Park', 'Park City', 'Park Forest', 'Park Ridge', 'Peotone', 'Phoenix', 'Pingree Grove', 'Pistakee Highlands', 'Plainfield', 'Plano', 'Plato Center', 'Plattville', 'Port Barrington', 'Posen', 'Prairie Grove', 'Prairie View', 'Prestbury', 'Preston Heights', 'Prospect Heights', 'Richmond', 'Richton Park', 'Ridgefield', 'Ridgewood', 'Ringwood', 'River Forest', 'River Grove', 'Riverdale', 'Riverside', 'Riverwoods', 'Robbins', 'Rockdale', 'Rolling Meadows', 'Romeoville', 'Rondout', 'Roselle', 'Rosemont', 'Round Lake', 'Round Lake Beach', 'Round Lake Heights', 'Round Lake Park', 'Russell', 'Saint Charles', 'Sauk Village', 'Schaumburg', 'Schiller Park', 'Shorewood', 'Skokie', 'Sleepy Hollow', 'South Barrington', 'South Chicago Heights', 'South Elgin', 'South Holland', 'Spring Grove', 'Steger', 'Stickney', 'Stone Park', 'Streamwood', 'Sugar Grove', 'Summit', 'Sunny Crest', 'Symerton', 'Third Lake', 'Thornton', 'Tinley Park', 'Tower Lake', 'Tower Lakes', 'Trout Valley', 'Union', 'University Park', 'Valley View', 'Venetian Village', 'Vernon Hills', 'Villa Park', 'Virgil', 'Volo', 'Wadsworth', 'Warrenville', 'Wauconda', 'Waukegan', 'Wayne', 'Wedges Corner', 'West Chicago', 'West Dundee', 'Westchester', 'Western Springs', 'Westmont', 'Wheaton', 'Wheeling', 'Wildwood', 'Williams Park', 'Willow Springs', 'Willowbrook', 'Wilmette', 'Wilmington', 'Wilson', 'Wilton Center', 'Winfield', 'Winnetka', 'Winthrop Harbor', 'Wonder Lake', 'Wood Dale', 'Woodridge', 'Woodstock', 'Worth', 'York Center', 'Yorkfield', 'Yorkville', 'Zion')) and `estates`.`case_county` in ('Cook', 'Cook county', 'DuPage', 'DuPage county', 'Kane', 'Kane county', 'Kendall', 'Kendall county', 'Lake', 'Lake county', 'McHenry', 'McHenry county', 'Will', 'Will county') and `estates`.`type_id` in ('5', '4', '8', '12', '2', '11', '1', '6', '14', '10', '3', '7', '13', '15', '21') and `liens`.`year_of_mortgage` between '1980' and '2021' and `liens`.`deleted_at` is null and `foreclosures`.`balance_due` between '50000' and '1950000' and `listings`.`type` != 'probate' and `listings`.`deleted_at` is null order by `proxy_auctions`.`last_status_update` desc").Limit(5).Scan(&user)

	core.Server.DB.Connections["main"].Where("id = ?", userId).Find(&user)
	//db := r.WithContext("db")(gorm.DB)
	render.JSON(w, r, user) // A chi router helper for serializing and returning json
}

//func getUsers(w http.ResponseWriter, r *http.Request) {
//	params := chi.URLParam(r, "params")
//	page := chi.URLParam(r, "page")
//	filters := chi.URLParam(r, "filters")
//	perPage := chi.URLParam(r, "perPage")
//	var Users []gormUser.User
//
//	// Create *DB object with our filters
//	filteredUsers := core.Server.DB.Connections["main"].Model(&Users)
//
//	// Declare variable for total rows
//	var totalRows int
//	// Populate total row count
//	filteredUsers.Count(&totalRows)
//
//	// Retrieve paginated rows
//	offset := perPage * page
//	filteredUsers.Limit(perPage).Offset().Find(&users)
//
//}

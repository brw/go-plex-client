package plex

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"
	"time"
)

// Plex contains fields that are required to make
// an api call to your plex server
type Plex struct {
	URL              string
	Token            string
	ClientIdentifier string
	Headers          headers
	HTTPClient       http.Client
	DownloadClient   http.Client
}

// SearchResults a list of media returned when searching
// for media via your plex server

// Provider ...
type Provider struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

// SearchMediaContainer ...
type SearchMediaContainer struct {
	MediaContainer
	Provider []Provider
}

// SearchResults ...
type SearchResults struct {
	MediaContainer SearchMediaContainer `json:"MediaContainer"`
}

// Metadata ...
type Metadata struct {
	Player                Player       `json:"Player"`
	Session               Session      `json:"Session"`
	User                  User         `json:"User"`
	AddedAt               int          `json:"addedAt"`
	Art                   string       `json:"art"`
	AudienceRating        float64      `json:"audienceRating"`
	AudienceRatingImage   string       `json:"audienceRatingImage"`
	ContentRating         string       `json:"contentRating"`
	Duration              int          `json:"duration"`
	GrandparentArt        string       `json:"grandparentArt"`
	GrandparentGUID       string       `json:"grandparentGuid"`
	GrandparentKey        string       `json:"grandparentKey"`
	GrandparentRatingKey  string       `json:"grandparentRatingKey"`
	GrandparentTheme      string       `json:"grandparentTheme"`
	GrandparentThumb      string       `json:"grandparentThumb"`
	GrandparentTitle      string       `json:"grandparentTitle"`
	GUID                  string       `json:"guid"`
	AltGUIDs              []AltGUID    `json:"Guid"`
	Index                 int64        `json:"index"`
	Key                   string       `json:"key"`
	LastViewedAt          int          `json:"lastViewedAt"`
	LibrarySectionID      int          `json:"librarySectionID"`
	LibrarySectionKey     string       `json:"librarySectionKey"`
	LibrarySectionTitle   string       `json:"librarySectionTitle"`
	Live                  string       `json:"live"`
	OriginallyAvailableAt string       `json:"originallyAvailableAt"`
	OriginalTitle         string       `json:"originalTitle"`
	ParentGUID            string       `json:"parentGuid"`
	ParentIndex           int64        `json:"parentIndex"`
	ParentKey             string       `json:"parentKey"`
	ParentRatingKey       string       `json:"parentRatingKey"`
	ParentThumb           string       `json:"parentThumb"`
	ParentTitle           string       `json:"parentTitle"`
	RatingCount           int          `json:"ratingCount"`
	Ratings               []Rating     `json:"Rating"`
	Rating                float64      `json:"rating"`
	RatingKey             string       `json:"ratingKey"`
	SessionKey            string       `json:"sessionKey"`
	Summary               string       `json:"summary"`
	Thumb                 string       `json:"thumb"`
	Media                 []Media      `json:"Media"`
	Title                 string       `json:"title"`
	TitleSort             string       `json:"titleSort"`
	Type                  string       `json:"type"`
	Slug                  string       `json:"slug"`
	ParentSlug            string       `json:"parentSlug"`
	GrandparentSlug       string       `json:"grandparentSlug"`
	UpdatedAt             int          `json:"updatedAt"`
	UserRating            float64      `json:"userRating,string"`
	ViewCount             json.Number  `json:"viewCount"`
	ViewOffset            int          `json:"viewOffset"`
	Year                  int          `json:"year"`
	Director              []TaggedData `json:"Director"`
	Writer                []TaggedData `json:"Writer"`
}

// AltGUID represents a Globally Unique Identifier for a metadata provider that is not actively being used.
type AltGUID struct {
	ID string `json:"id"`
}

// MetadataV1 ...
type MetadataV1 struct {
	Metadata
	Index            int64     `json:"index"`
	ParentIndex      int64     `json:"parentIndex"`
	AddedAt          int64     `json:"addedAt"`
	Duration         int64     `json:"duration"`
	LastViewedAt     Timestamp `json:"lastViewedAt"`
	LibrarySectionID string    `json:"librarySectionID"`
	Media            []MediaV1 `json:"Media"`
	Rating           float64   `json:"rating"`
	UpdatedAt        Timestamp `json:"updatedAt"`
	ViewOffset       int64     `json:"viewOffset"`
	Year             int       `json:"year"`
}

// Media media info
type Media struct {
	AspectRatio           string `json:"aspectRatio"`
	AudioChannels         int    `json:"audioChannels"`
	AudioCodec            string `json:"audioCodec"`
	AudioProfile          string `json:"audioProfile"`
	Bitrate               int    `json:"bitrate"`
	Container             string `json:"container"`
	DeletedAt             int    `json:"deletedAt"`
	Duration              int    `json:"duration"`
	Has64bitOffsets       bool   `json:"has64bitOffsets"`
	Height                int    `json:"height"`
	ID                    int    `json:"id"`
	OptimizedForStreaming int    `json:"optimizedForStreaming"`
	Selected              bool   `json:"selected"`
	VideoCodec            string `json:"videoCodec"`
	VideoFrameRate        string `json:"videoFrameRate"`
	VideoProfile          string `json:"videoProfile"`
	VideoResolution       string `json:"videoResolution"`
	Width                 int    `json:"width"`
	Part                  []Part `json:"Part"`
}

// MediaV1 media information version 1
type MediaV1 struct {
	Media
	Part                  []PartV1  `json:"Part"`
	AudioChannels         int       `json:"audioChannels"`
	AspectRatio           float32   `json:"aspectRatio,string"`
	Bitrate               int       `json:"bitrate"`
	Duration              Timestamp `json:"duration"`
	Has64bitOffsets       bool      `json:"has64bitOffsets"`
	Height                int       `json:"height"`
	ID                    int       `json:"id,string"`
	OptimizedForStreaming bool      `json:"optimizedForStreaming"`
	Width                 int       `json:"width"`
}

// MediaContainer contains media info
type MediaContainer struct {
	Metadata            []Metadata `json:"Metadata"`
	AllowSync           bool       `json:"allowSync"`
	Identifier          string     `json:"identifier"`
	LibrarySectionID    int        `json:"librarySectionID"`
	LibrarySectionTitle string     `json:"librarySectionTitle"`
	LibrarySectionUUID  string     `json:"librarySectionUUID"`
	MediaTagPrefix      string     `json:"mediaTagPrefix"`
	MediaTagVersion     int        `json:"mediaTagVersion"`
	Size                int        `json:"size"`
}

// MediaMetadata ...
type MediaMetadata struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
}

// Location is the path of a plex server directory
type Location struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

// Directory shows plex directory metadata
type Directory struct {
	Location   []Location `json:"Location"`
	Agent      string     `json:"agent"`
	AllowSync  bool       `json:"allowSync"`
	Art        string     `json:"art"`
	Composite  string     `json:"composite"`
	CreatedAt  int        `json:"createdAt"`
	Filter     bool       `json:"filters"`
	Key        string     `json:"key"`
	Language   string     `json:"language"`
	Refreshing bool       `json:"refreshing"`
	Scanner    string     `json:"scanner"`
	Thumb      string     `json:"thumb"`
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	UpdatedAt  int        `json:"updatedAt"`
	UUID       string     `json:"uuid"`
}

// LibrarySections metadata of your library contents
type LibrarySections struct {
	MediaContainer struct {
		Directory []Directory `json:"Directory"`
	} `json:"MediaContainer"`
}

// TaggedData ...
type TaggedData struct {
	Tag    string      `json:"tag"`
	Filter string      `json:"filter"`
	ID     json.Number `json:"id"`
}

// Role ...
type Role struct {
	TaggedData
	Role  string `json:"role"`
	Thumb string `json:"thumb"`
}

// MetadataChildren returns metadata about a piece of media (tv show, movie, music, etc)
type MetadataChildren struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
}

// SearchResultsEpisode contains metadata about an episode
type SearchResultsEpisode struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
}

type plexResponse struct {
	Children []struct {
		ElementType string `json:"_elementType"`
		Count       string `json:"count"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
	ElementType                   string `json:"_elementType"`
	AllowCameraUpload             string `json:"allowCameraUpload"`
	AllowChannelAccess            string `json:"allowChannelAccess"`
	AllowSync                     string `json:"allowSync"`
	BackgroundProcessing          string `json:"backgroundProcessing"`
	Certificate                   string `json:"certificate"`
	CompanionProxy                string `json:"companionProxy"`
	FriendlyName                  string `json:"friendlyName"`
	HubSearch                     string `json:"hubSearch"`
	MachineIdentifier             string `json:"machineIdentifier"`
	Multiuser                     string `json:"multiuser"`
	MyPlex                        string `json:"myPlex"`
	MyPlexMappingState            string `json:"myPlexMappingState"`
	MyPlexSigninState             string `json:"myPlexSigninState"`
	MyPlexSubscription            string `json:"myPlexSubscription"`
	MyPlexUsername                string `json:"myPlexUsername"`
	Platform                      string `json:"platform"`
	PlatformVersion               string `json:"platformVersion"`
	RequestParametersInCookie     string `json:"requestParametersInCookie"`
	Sync                          string `json:"sync"`
	TranscoderActiveVideoSessions string `json:"transcoderActiveVideoSessions"`
	TranscoderAudio               string `json:"transcoderAudio"`
	TranscoderLyrics              string `json:"transcoderLyrics"`
	TranscoderPhoto               string `json:"transcoderPhoto"`
	TranscoderSubtitles           string `json:"transcoderSubtitles"`
	TranscoderVideo               string `json:"transcoderVideo"`
	TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
	TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
	TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
	UpdatedAt                     string `json:"updatedAt"`
	Version                       string `json:"version"`
}

type killTranscodeResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

// CreateLibraryParams params required to create a library
type CreateLibraryParams struct {
	Name        string
	Location    string
	LibraryType string
	Agent       string
	Scanner     string
	Language    string
}

// DevicesResponse  metadata of a device that has connected to your server
type DevicesResponse struct {
	ID         int    `json:"id"`
	LastSeenAt string `json:"lastSeenAt"`
	Name       string `json:"name"`
	Product    string `json:"product"`
	Version    string `json:"version"`
}

// Friends are the plex accounts that have access to your server
type Friends struct {
	ID                        int    `xml:"id,attr"`
	Title                     string `xml:"title,attr"`
	Thumb                     string `xml:"thumb,attr"`
	Protected                 string `xml:"protected,attr"`
	Home                      string `xml:"home,attr"`
	AllowSync                 string `xml:"allowSync,attr"`
	AllowCameraUpload         string `xml:"allowCameraUpload,attr"`
	AllowChannels             string `xml:"allowChannels,attr"`
	FilterAll                 string `xml:"filterAll,attr"`
	FilterMovies              string `xml:"filterMovies,attr"`
	FilterMusic               string `xml:"filterMusic,attr"`
	FilterPhotos              string `xml:"filterPhotos,attr"`
	FilterTelevision          string `xml:"filterTelevision,attr"`
	Restricted                string `xml:"restricted,attr"`
	Username                  string `xml:"username,attr"`
	Email                     string `xml:"email,attr"`
	RecommendationsPlaylistID string `xml:"recommendationsPlaylistId,attr"`
	Server                    struct {
		ID                string `xml:"id,attr"`
		ServerID          string `xml:"serverId,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		Name              string `xml:"name,attr"`
		LastSeenAt        string `xml:"lastSeenAt,attr"`
		NumLibraries      string `xml:"numLibraries,attr"`
		AllLibraries      string `xml:"allLibraries,attr"`
		Owned             string `xml:"owned,attr"`
		Pending           string `xml:"pending,attr"`
	} `xml:"Server"`
}

type friendsResponse struct {
	XMLName           xml.Name  `xml:"MediaContainer"`
	FriendlyName      string    `xml:"friendlyName,attr"`
	Identifier        string    `xml:"identifier,attr"`
	MachineIdentifier string    `xml:"machineIdentifier,attr"`
	TotalSize         string    `xml:"totalSize,attr"`
	Size              int       `xml:"size,attr"`
	User              []Friends `xml:"User"`
}

type resultResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Response struct {
		Code   int    `xml:"code,attr"`
		Status string `xml:"status,attr"`
	} `xml:"Response"`
}

type inviteFriendResponse struct {
	ID                json.Number `json:"id"`
	Name              string      `json:"name"`
	OwnerID           json.Number `json:"ownerId"`
	InvitedID         json.Number `json:"invitedId"`
	InvitedEmail      string      `json:"invitedEmail"`
	ServerID          json.Number `json:"serverId"`
	Accepted          bool        `json:"accepted"`
	AcceptedAt        string      `json:"acceptedAt"`
	DeletedAt         string      `json:"deletedAt"`
	LeftAt            string      `json:"leftAt"`
	Owned             bool        `json:"owned"`
	InviteToken       string      `json:"inviteToken"`
	MachineIdentifier string      `json:"machineIdentifier"`
	LastSeenAt        time.Time   `json:"lastSeenAt"`
	NumLibraries      json.Number `json:"numLibraries"`
	Invited           struct {
		ID         json.Number `json:"id"`
		UUID       string      `json:"uuid"`
		Title      string      `json:"title"`
		Username   string      `json:"username"`
		Restricted bool        `json:"restricted"`
		Thumb      string      `json:"thumb"`
		Status     string      `json:"status"`
	} `json:"invited"`
	SharingSettings struct {
		AllowChannels    bool   `json:"allowChannels"`
		FilterMovies     string `json:"filterMovies"`
		FilterMusic      string `json:"filterMusic"`
		FilterPhotos     string `json:"filterPhotos"`
		FilterTelevision string `json:"filterTelevision"`
		// FilterAll ??? I get null when testing. idk the true type
		FilterAll          interface{} `json:"filterAll"`
		AllowSync          bool        `json:"allowSync"`
		AllowCameraUpload  bool        `json:"allowCameraUpload"`
		AllowSubtitleAdmin bool        `json:"allowSubtitleAdmin"`
		AllowTuners        json.Number `json:"allowTuners"`
	} `json:"sharingSettings"`
	Libraries []struct {
		ID    json.Number `json:"id"`
		Key   json.Number `json:"key"`
		Title string      `json:"title"`
		Type  string      `json:"type"`
	} `json:"libraries"`
	AllLibraries bool `json:"allLibraries"`
}

// InviteFriendParams are the params to invite a friend
type InviteFriendParams struct {
	UsernameOrEmail string
	MachineID       string
	Label           string
	LibraryIDs      []int
}

// UpdateFriendParams optional parameters to update your friends access to your server
type UpdateFriendParams struct {
	AllowSync         string
	AllowCameraUpload string
	AllowChannels     string
	FilterMovies      string
	FilterTelevision  string
	FilterMusic       string
	FilterPhotos      string
}
type inviteFriendBody struct {
	InvitedEmail      string               `json:"invitedEmail"`
	LibrarySectionIDs []int                `json:"librarySectionIds"`
	MachineIdentifier string               `json:"machineIdentifier"`
	Settings          inviteFriendSettings `json:"settings"`
}

type inviteFriendSettings struct {
	AllowCameraUpload string `json:"allowCameraUpload"`
	AllowSync         string `json:"allowSync"`
	FilterMovies      string `json:"filterMovies"`
	FilterMusic       string `json:"filterMusic"`
	FilterTelevision  string `json:"filterTelevision"`
}

type resourcesResponse struct {
	XMLName xml.Name     `xml:"MediaContainer"`
	Size    int          `xml:"size,attr"`
	Device  []PMSDevices `xml:"Device"`
}

type terminateSessionResponse struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Size    int      `xml:"size,attr"`
}

// PMSDevices is the result of the https://plex.tv/pms/resources endpoint
type PMSDevices struct {
	Name                 string       `json:"name" xml:"name,attr"`
	Product              string       `json:"product" xml:"product,attr"`
	ProductVersion       string       `json:"productVersion" xml:"productVersion,attr"`
	Platform             string       `json:"platform" xml:"platform,attr"`
	PlatformVersion      string       `json:"platformVersion" xml:"platformVersion,attr"`
	Device               string       `json:"device" xml:"device,attr"`
	ClientIdentifier     string       `json:"clientIdentifier" xml:"clientIdentifier,attr"`
	CreatedAt            string       `json:"createdAt" xml:"createdAt,attr"`
	LastSeenAt           string       `json:"lastSeenAt" xml:"lastSeenAt,attr"`
	Provides             string       `json:"provides" xml:"provides,attr"`
	Owned                string       `json:"owned" xml:"owned,attr"`
	AccessToken          string       `json:"accessToken" xml:"accessToken,attr"`
	HTTPSRequired        int          `json:"httpsRequired" xml:"httpsRequired,attr"`
	Synced               string       `json:"synced" xml:"synced,attr"`
	Relay                int          `json:"relay" xml:"relay,attr"`
	PublicAddressMatches string       `json:"publicAddressMatches" xml:"publicAddressMatches,attr"`
	PublicAddress        string       `json:"publicAddress" xml:"publicAddress,attr"`
	Presence             string       `json:"presence" xml:"presence,attr"`
	Connection           []Connection `json:"connection" xml:"Connection"`
}

// Connection lists options to connect to a device
type Connection struct {
	Protocol string `json:"protocol" xml:"protocol,attr"`
	Address  string `json:"address" xml:"address,attr"`
	Port     string `json:"port" xml:"port,attr"`
	URI      string `json:"uri" xml:"uri,attr"`
	Local    int    `json:"local" xml:"local,attr"`
	Relay    bool   `json:"relay" xml:"relay,attr"`
}

// BaseAPIResponse info about the Plex Media Server
type BaseAPIResponse struct {
	MediaContainer struct {
		Directory []struct {
			Count int64  `json:"count"`
			Key   string `json:"key"`
			Title string `json:"title"`
		} `json:"Directory"`
		AllowCameraUpload             bool   `json:"allowCameraUpload"`
		AllowChannelAccess            bool   `json:"allowChannelAccess"`
		AllowSharing                  bool   `json:"allowSharing"`
		AllowSync                     bool   `json:"allowSync"`
		BackgroundProcessing          bool   `json:"backgroundProcessing"`
		Certificate                   bool   `json:"certificate"`
		CompanionProxy                bool   `json:"companionProxy"`
		CountryCode                   string `json:"countryCode"`
		Diagnostics                   string `json:"diagnostics"`
		EventStream                   bool   `json:"eventStream"`
		FriendlyName                  string `json:"friendlyName"`
		HubSearch                     bool   `json:"hubSearch"`
		ItemClusters                  bool   `json:"itemClusters"`
		Livetv                        int64  `json:"livetv"`
		MachineIdentifier             string `json:"machineIdentifier"`
		MediaProviders                bool   `json:"mediaProviders"`
		Multiuser                     bool   `json:"multiuser"`
		MyPlex                        bool   `json:"myPlex"`
		MyPlexMappingState            string `json:"myPlexMappingState"`
		MyPlexSigninState             string `json:"myPlexSigninState"`
		MyPlexSubscription            bool   `json:"myPlexSubscription"`
		MyPlexUsername                string `json:"myPlexUsername"`
		OwnerFeatures                 string `json:"ownerFeatures"`
		PhotoAutoTag                  bool   `json:"photoAutoTag"`
		Platform                      string `json:"platform"`
		PlatformVersion               string `json:"platformVersion"`
		PluginHost                    bool   `json:"pluginHost"`
		ReadOnlyLibraries             bool   `json:"readOnlyLibraries"`
		RequestParametersInCookie     bool   `json:"requestParametersInCookie"`
		Size                          int64  `json:"size"`
		StreamingBrainABRVersion      int64  `json:"streamingBrainABRVersion"`
		StreamingBrainVersion         int64  `json:"streamingBrainVersion"`
		Sync                          bool   `json:"sync"`
		TranscoderActiveVideoSessions int64  `json:"transcoderActiveVideoSessions"`
		TranscoderAudio               bool   `json:"transcoderAudio"`
		TranscoderLyrics              bool   `json:"transcoderLyrics"`
		TranscoderPhoto               bool   `json:"transcoderPhoto"`
		TranscoderSubtitles           bool   `json:"transcoderSubtitles"`
		TranscoderVideo               bool   `json:"transcoderVideo"`
		TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
		TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
		TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
		UpdatedAt                     int64  `json:"updatedAt"`
		Updater                       bool   `json:"updater"`
		Version                       string `json:"version"`
		VoiceSearch                   bool   `json:"voiceSearch"`
	} `json:"MediaContainer"`
}

// UserPlexTV plex.tv user. should be used when interacting with plex.tv as the id is an int
type UserPlexTV struct {
	XMLName xml.Name `xml:"user"`
	// ID is an int when signing in to Plex.tv but a string when access own server
	ID                  int    `json:"id" xml:"id,attr"`
	UUID                string `json:"uuid" xml:"uuid,attr"`
	Email               string `json:"email" xml:"email,attr"`
	JoinedAt            string `json:"joined_at" xml:"joined-at"`
	Username            string `json:"username" xml:"username,attr"`
	Thumb               string `json:"thumb" xml:"thumb,attr"`
	HasPassword         bool   `json:"hasPassword" xml:"hasPassword,attr"`
	AuthToken           string `json:"authToken" xml:"authToken,attr"`
	AuthenticationToken string `json:"authenticationToken" xml:"authentificationToken,attr"`
	Subscription        struct {
		Active  bool   `json:"active" xml:"active,attr"`
		Status  string `json:"Active" xml:"status,attr"`
		Plan    string `json:"lifetime" xml:"plan,attr"`
		Feature []struct {
			ID string `json:"id" xml:"id,attr"`
		} `json:"feature" xml:"feature"`
	} `json:"subscription" xml:"subscription"`
	Roles struct {
		Role []struct {
			ID string `json:"id" xml:"id,attr"`
		} `json:"roles" xml:"role"`
	} `json:"roles" xml:"roles"`
	Entitlements []struct {
		Entitlement []struct {
			ID string `json:"id" xml:"id,attr"`
		} `json:"entitlement" xml:"entitlement"`
	} `json:"entitlements" xml:"entitlements"`
	ConfirmedAt string      `json:"confirmedAt"`
	ForumID     json.Number `json:"forumId"`
	RememberMe  bool        `json:"rememberMe"`
	Title       string      `json:"title" xml:"title,attr"`
}

// User plex server user. only difference is id is a string
type User struct {
	// ID is an int when signing in to Plex.tv but a string when access own server
	ID                  string `json:"id"`
	UUID                string `json:"uuid"`
	Email               string `json:"email"`
	JoinedAt            string `json:"joined_at"`
	Username            string `json:"username"`
	Thumb               string `json:"thumb"`
	HasPassword         bool   `json:"hasPassword"`
	AuthToken           string `json:"authToken"`
	AuthenticationToken string `json:"authenticationToken"`
	Subscription        struct {
		Active   bool     `json:"active"`
		Status   string   `json:"Active"`
		Plan     string   `json:"lifetime"`
		Features []string `json:"features"`
	} `json:"subscription"`
	Roles struct {
		Roles []string `json:"roles"`
	} `json:"roles"`
	Entitlements []string `json:"entitlements"`
	ConfirmedAt  string   `json:"confirmedAt"`
	ForumID      string   `json:"forumId"`
	RememberMe   bool     `json:"rememberMe"`
	Title        string   `json:"title"`
}

// SignInResponse response from plex.tv sign in
type SignInResponse UserPlexTV

// ServerInfo is the result of the https://plex.tv/api/servers endpoint
type ServerInfo struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		AccessToken       string `xml:"accessToken,attr"`
		Name              string `xml:"name,attr"`
		Address           string `xml:"address,attr"`
		Port              string `xml:"port,attr"`
		Version           string `xml:"version,attr"`
		Scheme            string `xml:"scheme,attr"`
		Host              string `xml:"host,attr"`
		LocalAddresses    string `xml:"localAddresses,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		CreatedAt         string `xml:"createdAt,attr"`
		UpdatedAt         string `xml:"updatedAt,attr"`
		Owned             string `xml:"owned,attr"`
		Synced            string `xml:"synced,attr"`
	} `xml:"Server"`
}

// SectionIDResponse the section id (or library id) of your server
// useful when inviting a user to the server
type SectionIDResponse struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		Name              string           `xml:"name,attr"`
		Address           string           `xml:"address,attr"`
		Port              string           `xml:"port,attr"`
		Version           string           `xml:"version,attr"`
		Scheme            string           `xml:"scheme,attr"`
		Host              string           `xml:"host,attr"`
		LocalAddresses    string           `xml:"localAddresses,attr"`
		MachineIdentifier string           `xml:"machineIdentifier,attr"`
		CreatedAt         int              `xml:"createdAt,attr"`
		UpdatedAt         int              `xml:"updatedAt,attr"`
		Owned             int              `xml:"owned,attr"`
		Synced            string           `xml:"synced,attr"`
		Section           []ServerSections `xml:"Section"`
	} `xml:"Server"`
}

// ServerSections contains information of your library sections
type ServerSections struct {
	ID    int    `xml:"id,attr"`
	Key   string `xml:"key,attr"`
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
}

// LibraryLabels are the existing labels set on your server
type LibraryLabels struct {
	ElementType     string `json:"_elementType"`
	AllowSync       string `json:"allowSync"`
	Art             string `json:"art"`
	Content         string `json:"content"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
	Thumb           string `json:"thumb"`
	Title1          string `json:"title1"`
	Title2          string `json:"title2"`
	ViewGroup       string `json:"viewGroup"`
	ViewMode        string `json:"viewMode"`
	Children        []struct {
		ElementType string `json:"_elementType"`
		FastKey     string `json:"fastKey"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
}

type headers struct {
	Platform               string
	PlatformVersion        string
	Provides               string
	Product                string
	Version                string
	Device                 string
	ContainerSize          string
	ContainerStart         string
	Token                  string
	Accept                 string
	ContentType            string
	ClientIdentifier       string
	TargetClientIdentifier string
}

type request struct {
	headers
}

// Sessions

// TranscodeSessionsResponse is the result for transcode session endpoint /transcode/sessions
type TranscodeSessionsResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

// Rating ...
type Rating struct {
	Count int         `json:"count,string"`
	Image string      `json:"image"`
	Type  string      `json:"type"`
	Value FixedRating `json:"value"`
}

type FixedRating float32

func (value *FixedRating) UnmarshalJSON(data []byte) error {
	var isAlreadyFloat float32

	if err := json.Unmarshal(data, &isAlreadyFloat); err == nil {
		*value = FixedRating(isAlreadyFloat)
		return nil
	}

	var isString string

	if err := json.Unmarshal(data, &isString); err != nil {
		return err
	}

	conv, _ := strconv.ParseFloat(isString, 32)
	*value = FixedRating(conv)
	return nil
}

// Stream ...
type Stream struct {
	AlbumGain          string  `json:"albumGain"`
	AlbumPeak          string  `json:"albumPeak"`
	AlbumRange         string  `json:"albumRange"`
	Anamorphic         bool    `json:"anamorphic"`
	AudioChannelLayout string  `json:"audioChannelLayout"`
	BitDepth           int     `json:"bitDepth"`
	Bitrate            int     `json:"bitrate"`
	BitrateMode        string  `json:"bitrateMode"`
	Cabac              string  `json:"cabac"`
	Channels           int     `json:"channels"`
	ChromaLocation     string  `json:"chromaLocation"`
	ChromaSubsampling  string  `json:"chromaSubsampling"`
	Codec              string  `json:"codec"`
	CodecID            string  `json:"codecID"`
	ColorRange         string  `json:"colorRange"`
	ColorSpace         string  `json:"colorSpace"`
	Default            bool    `json:"default"`
	DisplayTitle       string  `json:"displayTitle"`
	Duration           float64 `json:"duration"`
	FrameRate          float64 `json:"frameRate"`
	FrameRateMode      string  `json:"frameRateMode"`
	Gain               string  `json:"gain"`
	HasScalingMatrix   bool    `json:"hasScalingMatrix"`
	Height             int     `json:"height"`
	ID                 int     `json:"id"`
	Index              int     `json:"index"`
	Language           string  `json:"language"`
	LanguageCode       string  `json:"languageCode"`
	Level              int     `json:"level"`
	Location           string  `json:"location"`
	Loudness           string  `json:"loudness"`
	Lra                string  `json:"lra"`
	Peak               string  `json:"peak"`
	PixelAspectRatio   string  `json:"pixelAspectRatio"`
	PixelFormat        string  `json:"pixelFormat"`
	Profile            string  `json:"profile"`
	RefFrames          int     `json:"refFrames"`
	SamplingRate       int     `json:"samplingRate"`
	ScanType           string  `json:"scanType"`
	Selected           bool    `json:"selected"`
	StreamIdentifier   string  `json:"streamIdentifier"`
	StreamType         int     `json:"streamType"`
	Width              int     `json:"width"`
}

// StreamV1 stream info version 1
type StreamV1 struct {
	Stream
	BitDepth         int     `json:"bitDepth"`
	Default          bool    `json:"default"`
	Bitrate          int     `json:"bitrate"`
	FrameRate        float64 `json:"frameRate"`
	HasScalingMatrix bool    `json:"hasScalingMatrix"`
	Height           int     `json:"height"`
	Width            int     `json:"width"`
	ID               int     `json:"id,string"`
	Index            int     `json:"index"`
	Level            int     `json:"level"`
	RefFrames        int     `json:"refFrames"`
	StreamType       int     `json:"streamType"`
	Channels         int     `json:"channels"`
	SamplingRate     int     `json:"samplingRate"`
	Selected         bool    `json:"selected"`
}

// Part ...
type Part struct {
	AudioProfile          string   `json:"audioProfile"`
	Container             string   `json:"container"`
	Decision              string   `json:"decision"`
	Duration              int64    `json:"duration"`
	File                  string   `json:"file"`
	Has64bitOffsets       bool     `json:"has64bitOffsets"`
	HasThumbnail          string   `json:"hasThumbnail"`
	ID                    int      `json:"id"`
	Key                   string   `json:"key"`
	OptimizedForStreaming bool     `json:"optimizedForStreaming"`
	Selected              bool     `json:"selected"`
	Size                  int      `json:"size"`
	Stream                []Stream `json:"Stream"`
	VideoProfile          string   `json:"videoProfile"`
}

// PartV1 part version 1
type PartV1 struct {
	Part
	Duration              int64      `json:"duration"`
	Has64bitOffsets       bool       `json:"has64bitOffsets"`
	ID                    int        `json:"id,string"`
	OptimizedForStreaming bool       `json:"optimizedForStreaming"`
	Size                  int        `json:"size"`
	Stream                []StreamV1 `json:"Stream"`
}

// Player ...
type Player struct {
	Address             string `json:"address"`
	Device              string `json:"device"`
	Local               bool   `json:"local"`
	MachineIdentifier   string `json:"machineIdentifier"`
	Model               string `json:"model"`
	Platform            string `json:"platform"`
	PlatformVersion     string `json:"platformVersion"`
	Product             string `json:"product"`
	Profile             string `json:"profile"`
	RemotePublicAddress string `json:"remotePublicAddress"`
	State               string `json:"state"`
	Title               string `json:"title"`
	UserID              int    `json:"userID"`
	Vendor              string `json:"vendor"`
	Version             string `json:"version"`
}

// Session ...
type Session struct {
	Bandwidth int    `json:"bandwidth"`
	ID        string `json:"id"`
	Location  string `json:"location"`
}

// CurrentSessions metadata of users consuming media
type CurrentSessions struct {
	MediaContainer struct {
		Metadata []MetadataV1 `json:"Metadata"`
		Size     int          `json:"size"`
	} `json:"MediaContainer"`
}

type DiscoverTaggedData struct {
	Tag    string `json:"tag"`
	Filter string `json:"filter"`
	ID     string `json:"id"`
}

type DiscoverRole struct {
	DiscoverTaggedData
	Key   string `json:"key"`
	Role  string `json:"role"`
	Slug  string `json:"slug"`
	Thumb string `json:"thumb"`
	Type  string `json:"type"`
}

type DiscoverGenre struct {
	DiscoverTaggedData
	Context   string `json:"context"`
	Key       string `json:"key"`
	RatingKey string `json:"ratingKey"`
	Slug      string `json:"slug"`
	Thumb     string `json:"thumb"`
	Type      string `json:"type"`
}

type DiscoverGuid struct {
	ID string `json:"id"`
}

type DiscoverImage struct {
	Alt  string `json:"alt"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type DiscoverRating struct {
	Image string  `json:"image"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type DiscoverSeasonType struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Tag    string `json:"tag"`
	Title  string `json:"title"`
}

type DiscoverSummary struct {
	Size int    `json:"size"`
	Tag  string `json:"tag"`
	Type string `json:"type"`
}

type DiscoverMetadata struct {
	Countries             []DiscoverTaggedData `json:"Country"`
	Directors             []DiscoverRole       `json:"Director"`
	Genres                []DiscoverGenre      `json:"Genre"`
	GUIDs                 []DiscoverGuid       `json:"Guid"`
	Images                []DiscoverImage      `json:"Image"`
	Networks              []DiscoverTaggedData `json:"Network"`
	Ratings               []DiscoverRating     `json:"Rating"`
	Roles                 []DiscoverRole       `json:"Role"`
	SeasonTypes           []DiscoverSeasonType `json:"SeasonType"`
	Studios               []DiscoverTaggedData `json:"Studio"`
	Summaries             []DiscoverSummary    `json:"Summary"`
	Writers               []DiscoverRole       `json:"Writer"`
	AddedAt               int                  `json:"addedAt"`
	Art                   string               `json:"art"`
	AudienceRating        float64              `json:"audienceRating"`
	AudienceRatingImage   string               `json:"audienceRatingImage"`
	Banner                string               `json:"banner"`
	ContentRating         string               `json:"contentRating"`
	Duration              int                  `json:"duration"`
	GUID                  string               `json:"guid"`
	IMDBRatingCount       int                  `json:"imdbRatingCount"`
	Key                   string               `json:"key"`
	LeafCount             int                  `json:"leafCount"`
	OriginalTitle         string               `json:"originalTitle"`
	OriginallyAvailableAt string               `json:"originallyAvailableAt"`
	PublicPagesURL        string               `json:"publicPagesURL"`
	Rating                float64              `json:"rating"`
	RatingImage           string               `json:"ratingImage"`
	RatingKey             string               `json:"ratingKey"`
	Slug                  string               `json:"slug"`
	Studio                string               `json:"studio"`
	Summary               string               `json:"summary"`
	Tagline               string               `json:"tagline"`
	Theme                 string               `json:"theme"`
	Thumb                 string               `json:"thumb"`
	Title                 string               `json:"title"`
	Type                  string               `json:"type"`
	UserState             bool                 `json:"userState"`
	Year                  int                  `json:"year"`
}

type DiscoverMediaContainer struct {
	Metadata            []DiscoverMetadata `json:"Metadata"`
	Identifier          string             `json:"identifier"`
	LibrarySectionID    string             `json:"librarySectionID"`
	LibrarySectionTitle string             `json:"librarySectionTitle"`
	Offset              int                `json:"offset"`
	Size                int                `json:"size"`
	TotalSize           int                `json:"totalSize"`
}

type DiscoverMetadataResponse struct {
	MediaContainer DiscoverMediaContainer `json:"MediaContainer"`
}

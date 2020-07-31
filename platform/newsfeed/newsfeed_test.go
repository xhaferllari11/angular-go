package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"Item","Item Body"})
	if len(feed.Items) == 0 {
		t.Errorf("Item not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{"Item","Item Body"})
	feed.Add(Item{"Item2","Another Body"})
	result := feed.GetAll()
	if len(result) != 2 {
		t.Errorf("Did not get all items")
	}
}
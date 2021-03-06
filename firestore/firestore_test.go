package firestore

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/keys-pub/keys"
	"github.com/keys-pub/keys/ds"
	"github.com/keys-pub/keys/util"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/option"
)

const testURL = "firestore://chilltest-3297b"

var ctx = context.TODO()

func testFirestore(t *testing.T, clear bool) *Firestore {
	opts := []option.ClientOption{option.WithCredentialsFile("credentials.json")}
	fs, err := NewFirestore(testURL, opts...)
	require.NoError(t, err)
	fs.Test = true
	require.NoError(t, err)
	if clear {
		_, err := fs.Delete(ctx, "/")
		require.NoError(t, err)
	}
	return fs
}

type clock struct {
	t time.Time
}

func newClock() *clock {
	t := util.TimeFromMillis(1234567890000)
	return &clock{
		t: t,
	}
}

func (c *clock) Now() time.Time {
	c.t = c.t.Add(time.Millisecond)
	return c.t
}

func TestFirestore(t *testing.T) {
	// SetContextLogger(NewContextLogger(DebugLevel))
	fs := testFirestore(t, true)
	testDocumentStore(t, fs)
}

func TestFirestorePath(t *testing.T) {
	fs := testFirestore(t, true)
	testDocumentStorePath(t, fs)
}

func TestFirestoreListOptions(t *testing.T) {
	fs := testFirestore(t, true)
	testDocumentStoreListOptions(t, fs)
}

func TestFirestoreMetadata(t *testing.T) {
	// SetContextLogger(NewContextLogger(DebugLevel))
	fs := testFirestore(t, true)
	testMetadata(t, fs)
}

func TestEmptyIterator(t *testing.T) {
	iter := &docsIterator{}
	doc, err := iter.Next()
	require.NoError(t, err)
	require.Nil(t, doc)
	iter.Release()
}

func testDocumentStore(t *testing.T, dst ds.DocumentStore) {
	ctx := context.TODO()

	for i := 10; i <= 30; i = i + 10 {
		p := ds.Path("test1", fmt.Sprintf("key%d", i))
		err := dst.Create(ctx, p, []byte(fmt.Sprintf("value%d", i)))
		require.NoError(t, err)
	}
	for i := 10; i <= 30; i = i + 10 {
		p := ds.Path("test0", fmt.Sprintf("key%d", i))
		err := dst.Create(ctx, p, []byte(fmt.Sprintf("value%d", i)))
		require.NoError(t, err)
	}

	iter, err := dst.Documents(ctx, "test0", nil)
	require.NoError(t, err)
	doc, err := iter.Next()
	require.NoError(t, err)
	require.Equal(t, "/test0/key10", doc.Path)
	require.Equal(t, "value10", string(doc.Data))
	iter.Release()

	ok, err := dst.Exists(ctx, "/test0/key10")
	require.NoError(t, err)
	require.True(t, ok)
	doc, err = dst.Get(ctx, "/test0/key10")
	require.NoError(t, err)
	require.NotNil(t, doc)
	require.Equal(t, "value10", string(doc.Data))

	err = dst.Create(ctx, "/test0/key10", []byte{})
	require.EqualError(t, err, "path already exists /test0/key10")
	err = dst.Set(ctx, "/test0/key10", []byte("overwrite"))
	require.NoError(t, err)
	err = dst.Create(ctx, "/test0/key10", []byte("overwrite"))
	require.EqualError(t, err, "path already exists /test0/key10")
	doc, err = dst.Get(ctx, "/test0/key10")
	require.NoError(t, err)
	require.NotNil(t, doc)
	require.Equal(t, "overwrite", string(doc.Data))

	out, err := dst.GetAll(ctx, []string{"/test0/key10", "/test0/key20"})
	require.NoError(t, err)
	require.Equal(t, 2, len(out))
	require.Equal(t, "/test0/key10", out[0].Path)
	require.Equal(t, "/test0/key20", out[1].Path)

	ok, err = dst.Delete(ctx, "/test1/key10")
	require.True(t, ok)
	require.NoError(t, err)
	ok, err = dst.Delete(ctx, "/test1/key10")
	require.False(t, ok)
	require.NoError(t, err)

	ok, err = dst.Exists(ctx, "/test1/key10")
	require.NoError(t, err)
	require.False(t, ok)

	expected := `/test0/key10 overwrite
/test0/key20 value20
/test0/key30 value30
`

	var b bytes.Buffer
	iter, err = dst.Documents(context.TODO(), "test0", nil)
	require.NoError(t, err)
	err = ds.SpewOut(iter, nil, &b)
	require.NoError(t, err)
	require.Equal(t, expected, b.String())
	iter.Release()

	iter, err = dst.Documents(context.TODO(), "test0", nil)
	require.NoError(t, err)
	spew, err := ds.Spew(iter, nil)
	require.NoError(t, err)
	require.Equal(t, b.String(), spew.String())
	require.Equal(t, expected, spew.String())
	iter.Release()

	iter, err = dst.Documents(context.TODO(), "test0", &ds.DocumentsOpts{Prefix: "key1", PathOnly: true})
	require.NoError(t, err)
	doc, err = iter.Next()
	require.NoError(t, err)
	require.Equal(t, "/test0/key10", doc.Path)
	doc, err = iter.Next()
	require.NoError(t, err)
	require.Nil(t, doc)
	iter.Release()

	err = dst.Create(ctx, "", []byte{})
	require.EqualError(t, err, "invalid path /")
	err = dst.Set(ctx, "", []byte{})
	require.EqualError(t, err, "invalid path /")

	citer, err := dst.Collections(ctx, "")
	require.NoError(t, err)
	col, err := citer.Next()
	require.NoError(t, err)
	require.Equal(t, "/test0", col.Path)
	col, err = citer.Next()
	require.NoError(t, err)
	require.Equal(t, "/test1", col.Path)
	col, err = citer.Next()
	require.NoError(t, err)
	require.Nil(t, col)
	citer.Release()

	_, err = dst.Collections(ctx, "/test0")
	require.EqualError(t, err, "only root collections supported")
}

func testDocumentStorePath(t *testing.T, dst ds.DocumentStore) {
	ctx := context.TODO()

	err := dst.Create(ctx, "test/1", []byte("value1"))
	require.NoError(t, err)

	doc, err := dst.Get(ctx, "/test/1")
	require.NoError(t, err)
	require.NotNil(t, doc)

	ok, err := dst.Exists(ctx, "/test/1")
	require.NoError(t, err)
	require.True(t, ok)
	ok, err = dst.Exists(ctx, "test/1")
	require.NoError(t, err)
	require.True(t, ok)
}

func testDocumentStoreListOptions(t *testing.T, dst ds.DocumentStore) {
	ctx := context.TODO()

	err := dst.Create(ctx, "/test/1", []byte("val1"))
	require.NoError(t, err)
	err = dst.Create(ctx, "/test/2", []byte("val2"))
	require.NoError(t, err)
	err = dst.Create(ctx, "/test/3", []byte("val3"))
	require.NoError(t, err)

	for i := 1; i < 3; i++ {
		err := dst.Create(ctx, ds.Path("a", fmt.Sprintf("e%d", i)), []byte("🤓"))
		require.NoError(t, err)
	}
	for i := 1; i < 3; i++ {
		err := dst.Create(ctx, ds.Path("b", fmt.Sprintf("ea%d", i)), []byte("😎"))
		require.NoError(t, err)
	}
	for i := 1; i < 3; i++ {
		err := dst.Create(ctx, ds.Path("b", fmt.Sprintf("eb%d", i)), []byte("😎"))
		require.NoError(t, err)
	}
	for i := 1; i < 3; i++ {
		err := dst.Create(ctx, ds.Path("b", fmt.Sprintf("ec%d", i)), []byte("😎"))
		require.NoError(t, err)
	}
	for i := 1; i < 3; i++ {
		err := dst.Create(ctx, ds.Path("c", fmt.Sprintf("e%d", i)), []byte("😎"))
		require.NoError(t, err)
	}

	iter, err := dst.Documents(ctx, "test", nil)
	require.NoError(t, err)
	paths := []string{}
	for {
		doc, err := iter.Next()
		require.NoError(t, err)
		if doc == nil {
			break
		}
		paths = append(paths, doc.Path)
	}
	require.Equal(t, []string{"/test/1", "/test/2", "/test/3"}, paths)
	iter.Release()

	iter, err = dst.Documents(context.TODO(), "test", nil)
	require.NoError(t, err)
	b, err := ds.Spew(iter, nil)
	require.NoError(t, err)
	expected := `/test/1 val1
/test/2 val2
/test/3 val3
`
	require.Equal(t, expected, b.String())
	iter.Release()

	iter, err = dst.Documents(ctx, "b", &ds.DocumentsOpts{Prefix: "eb"})
	require.NoError(t, err)
	paths = []string{}
	for {
		doc, err := iter.Next()
		require.NoError(t, err)
		if doc == nil {
			break
		}
		paths = append(paths, doc.Path)
	}
	iter.Release()
	require.Equal(t, []string{"/b/eb1", "/b/eb2"}, paths)
}

func testMetadata(t *testing.T, dst ds.DocumentStore) {
	ctx := context.TODO()

	err := dst.Create(ctx, "/test/key1", []byte("value1"))
	require.NoError(t, err)

	doc, err := dst.Get(ctx, "/test/key1")
	require.NoError(t, err)
	require.NotNil(t, doc)
	createTime := util.TimeToMillis(doc.CreatedAt)
	require.True(t, createTime > 0)

	err = dst.Set(ctx, "/test/key1", []byte("value1b"))
	require.NoError(t, err)

	doc, err = dst.Get(ctx, "/test/key1")
	require.NoError(t, err)
	require.NotNil(t, doc)
	require.Equal(t, createTime, util.TimeToMillis(doc.CreatedAt))
	require.True(t, util.TimeToMillis(doc.UpdatedAt) > createTime)
}

func TestSigchains(t *testing.T) {
	clock := newClock()
	fs := testFirestore(t, true)
	scs := keys.NewSigchainStore(fs)

	kids := []keys.ID{}
	for i := 0; i < 6; i++ {
		key := keys.GenerateEdX25519Key()
		sc := keys.NewSigchain(key.ID())
		st, err := keys.NewSigchainStatement(sc, []byte("test"), key, "", clock.Now())
		require.NoError(t, err)
		err = sc.Add(st)
		require.NoError(t, err)
		err = scs.SaveSigchain(sc)
		require.NoError(t, err)
		kids = append(kids, key.ID())
	}

	sc, err := scs.Sigchain(kids[0])
	require.NoError(t, err)
	require.Equal(t, kids[0], sc.KID())
}

func ExampleNewFirestore() {
	url := "firestore://chilltest-3297b"
	opts := []option.ClientOption{option.WithCredentialsFile("credentials.json")}
	fs, err := NewFirestore(url, opts...)
	if err != nil {
		log.Fatal(err)
	}

	exists, err := fs.Exists(context.TODO(), "/test/1")
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		if _, err := fs.Delete(context.TODO(), "/test/1"); err != nil {
			log.Fatal(err)
		}
	}

	if err := fs.Create(context.TODO(), "/test/1", []byte("value1")); err != nil {
		log.Fatal(err)
	}

	entry, err := fs.Get(context.TODO(), "/test/1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", entry.Path)
	fmt.Printf("%s\n", entry.Data)
	// Output:
	// /test/1
	// value1
}

func TestDeleteAll(t *testing.T) {
	fs := testFirestore(t, true)

	err := fs.Set(context.TODO(), "/test/key1", []byte("val1"))
	require.NoError(t, err)
	err = fs.Set(context.TODO(), "/test/key2", []byte("val2"))
	require.NoError(t, err)

	err = fs.DeleteAll(context.TODO(), []string{"/test/key1", "/test/key2", "/test/key3"})
	require.NoError(t, err)

	doc, err := fs.Get(context.TODO(), "/test/key1")
	require.NoError(t, err)
	require.Nil(t, doc)
	doc, err = fs.Get(context.TODO(), "/test/key2")
	require.NoError(t, err)
	require.Nil(t, doc)
}

package mediator

import "sync"

type stationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []Itrain
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) requestArrival(t Itrain) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	//platform is not free so put the train in queue
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		//first train waiting gets to land in platform
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

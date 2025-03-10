package spectest

import (
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/commit"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/decided"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/futuremsg"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/latemsg"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/processmsg"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/messages"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/prepare"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/proposal"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/proposer"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/roundchange"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/startinstance"
	"testing"
)

type SpecTest interface {
	TestName() string
	Run(t *testing.T)
}

var AllTests = []SpecTest{
	//timeout.FirstRound(),
	//timeout.Round1(),
	//timeout.Round2(),
	//timeout.Round3(),
	//timeout.Round4(),
	//timeout.Round5(),
	//timeout.Round10(),
	//timeout.Round20(),
	//timeout.RoundTimeout(),
	//timeout.ErrorOnBroadcast(),
	//timeout.ErrorOnCreateMsg(),

	decided.HasQuorum(),
	decided.NoQuorum(),
	decided.DuplicateMsg(),
	decided.DuplicateSigners(),
	decided.ImparsableData(),
	decided.Invalid(),
	decided.InvalidData(),
	decided.InvalidValCheckData(),
	decided.LateDecided(),
	decided.LateDecidedBiggerQuorum(),
	decided.LateDecidedSmallerQuorum(),
	decided.PastInstance(),
	decided.UnknownSigner(),
	decided.WrongMsgType(),
	decided.WrongSignature(),
	decided.MultiDecidedInstances(),
	decided.FutureInstance(),
	decided.CurrentInstance(),
	decided.CurrentInstancePastRound(),
	decided.CurrentInstanceFutureRound(),

	processmsg.MsgError(),
	processmsg.SavedAndBroadcastedDecided(),
	processmsg.SingleConsensusMsg(),
	processmsg.FullDecided(),
	processmsg.InvalidIdentifier(),
	processmsg.NoInstanceRunning(),

	latemsg.LateCommit(),
	latemsg.LateCommitPastRound(),
	latemsg.LateCommitNoInstance(),
	latemsg.LateCommitPastInstance(),
	latemsg.LatePrepare(),
	latemsg.LatePrepareNoInstance(),
	latemsg.LatePreparePastInstance(),
	latemsg.LatePreparePastRound(),
	latemsg.LateProposal(),
	latemsg.LateProposalNoInstance(),
	latemsg.LateProposalPastInstance(),
	latemsg.LateProposalPastRound(),
	latemsg.LateRoundChange(),
	latemsg.LateRoundChangeNoInstance(),
	latemsg.LateRoundChangePastInstance(),
	latemsg.LateRoundChangePastRound(),
	latemsg.FullFlowAfterDecided(),

	futuremsg.NoSigners(),
	futuremsg.MultiSigners(),
	futuremsg.Cleanup(),
	futuremsg.DuplicateSigner(),
	futuremsg.F1FutureMsgs(),
	futuremsg.InvalidMsg(),
	futuremsg.UnknownSigner(),
	futuremsg.WrongSig(),

	startinstance.PostFutureDecided(),
	startinstance.FirstHeight(),
	startinstance.PreviousDecided(),
	startinstance.PreviousNotDecided(),
	startinstance.InvalidValue(),

	proposer.FourOperators(),
	proposer.SevenOperators(),
	proposer.TenOperators(),
	proposer.ThirteenOperators(),

	messages.RoundChangeDataInvalidJustifications(),
	messages.RoundChangeDataInvalidPreparedRound(),
	messages.RoundChangeDataInvalidPreparedValue(),
	messages.RoundChangePrePreparedJustifications(),
	messages.RoundChangeNotPreparedJustifications(),
	messages.CommitDataEncoding(),
	messages.MsgNilIdentifier(),
	messages.MsgNonZeroIdentifier(),
	messages.MsgTypeUnknown(),
	messages.PrepareDataEncoding(),
	messages.ProposeDataEncoding(),
	messages.MsgDataNil(),
	messages.MsgDataNonZero(),
	messages.SignedMsgSigTooShort(),
	messages.SignedMsgSigTooLong(),
	messages.SignedMsgNoSigners(),
	messages.SignedMsgDuplicateSigners(),
	messages.SignedMsgMultiSigners(),
	messages.GetRoot(),
	messages.SignedMessageEncoding(),
	messages.CreateProposal(),
	messages.CreateProposalPreviouslyPrepared(),
	messages.CreateProposalNotPreviouslyPrepared(),
	messages.CreatePrepare(),
	messages.CreateCommit(),
	messages.CreateRoundChange(),
	messages.CreateRoundChangePreviouslyPrepared(),
	messages.RoundChangeDataEncoding(),
	messages.PrepareDataInvalid(),
	messages.CommitDataInvalid(),
	messages.ProposalDataInvalid(),
	messages.SignedMessageSigner0(),

	tests.HappyFlow(),
	tests.SevenOperators(),
	tests.TenOperators(),
	tests.ThirteenOperators(),

	proposal.CurrentRoundPrevNotPrepared(),
	proposal.CurrentRoundPrevPrepared(),
	proposal.PastRoundProposalPrevPrepared(),
	proposal.NotPreparedPreviouslyJustification(),
	proposal.PreparedPreviouslyJustification(),
	proposal.DifferentJustifications(),
	proposal.JustificationsNotHeighest(),
	proposal.JustificationsValueNotJustified(),
	proposal.DuplicateMsg(),
	proposal.DuplicateMsgDifferentValue(),
	proposal.FirstRoundJustification(),
	proposal.FutureRoundPrevNotPrepared(),
	proposal.FutureRound(),
	proposal.ImparsableProposalData(),
	proposal.InvalidRoundChangeJustificationPrepared(),
	proposal.InvalidRoundChangeJustification(),
	proposal.PreparedPreviouslyNoRCJustificationQuorum(),
	proposal.NoRCJustification(),
	proposal.PreparedPreviouslyNoPrepareJustificationQuorum(),
	proposal.PreparedPreviouslyDuplicatePrepareMsg(),
	proposal.PreparedPreviouslyDuplicatePrepareQuorum(),
	proposal.PreparedPreviouslyDuplicateRCMsg(),
	proposal.PreparedPreviouslyDuplicateRCQuorum(),
	proposal.DuplicateRCMsg(),
	proposal.InvalidPrepareJustificationValue(),
	proposal.InvalidPrepareJustificationRound(),
	proposal.InvalidProposalData(),
	proposal.InvalidValueCheck(),
	proposal.MultiSigner(),
	proposal.PostDecided(),
	proposal.PostPrepared(),
	proposal.SecondProposalForRound(),
	proposal.WrongHeight(),
	proposal.WrongProposer(),
	proposal.WrongSignature(),
	proposal.UnknownSigner(),

	prepare.DuplicateMsg(),
	prepare.HappyFlow(),
	prepare.ImparsableProposalData(),
	prepare.InvalidPrepareData(),
	prepare.MultiSigner(),
	prepare.NoPreviousProposal(),
	prepare.OldRound(),
	prepare.FutureRound(),
	prepare.PostDecided(),
	prepare.WrongData(),
	prepare.WrongHeight(),
	prepare.WrongSignature(),
	prepare.UnknownSigner(),

	commit.CurrentRound(),
	commit.FutureRound(),
	commit.PastRound(),
	commit.DuplicateMsg(),
	commit.HappyFlow(),
	commit.InvalidCommitData(),
	commit.PostDecided(),
	commit.WrongData1(),
	commit.WrongData2(),
	commit.MultiSignerWithOverlap(),
	commit.MultiSignerNoOverlap(),
	commit.DuplicateSigners(),
	commit.NoPrevAcceptedProposal(),
	commit.WrongHeight(),
	commit.ImparsableCommitData(),
	commit.WrongSignature(),
	commit.UnknownSigner(),
	commit.InvalidValCheck(),
	commit.NoPrepareQuorum(),

	roundchange.HappyFlow(),
	roundchange.WrongHeight(),
	roundchange.WrongSig(),
	roundchange.UnknownSigner(),
	roundchange.MultiSigner(),
	roundchange.QuorumNotPrepared(),
	roundchange.QuorumPrepared(),
	roundchange.PeerPrepared(),
	roundchange.PeerPreparedDifferentHeights(),
	roundchange.JustificationWrongValue(),
	roundchange.JustificationWrongRound(),
	roundchange.JustificationNoQuorum(),
	roundchange.JustificationMultiSigners(),
	roundchange.JustificationInvalidSig(),
	roundchange.JustificationInvalidRound(),
	roundchange.JustificationInvalidPrepareData(),
	roundchange.JustificationDuplicateMsg(),
	roundchange.InvalidRoundChangeData(),
	roundchange.F1DifferentFutureRounds(),
	roundchange.F1DifferentFutureRoundsNotPrepared(),
	roundchange.PastRound(),
	roundchange.DuplicateMsgQuorum(),
	roundchange.DuplicateMsgQuorumPreparedRCFirst(),
	roundchange.DuplicateMsg(),
	roundchange.ImparsableRoundChangeData(),
	roundchange.NotProposer(),
	roundchange.ValidJustification(),
	roundchange.F1DuplicateSigner(),
	roundchange.F1DuplicateSignerNotPrepared(),
	roundchange.F1Speedup(),
	roundchange.F1SpeedupPrevPrepared(),
	roundchange.AfterProposal(),
	roundchange.RoundChangePartialQuorum(),
	roundchange.QuorumOrder2(),
	roundchange.QuorumOrder1(),
	roundchange.QuorumMsgNotPrepared(),
}
